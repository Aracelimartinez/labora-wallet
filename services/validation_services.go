package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labora-wallet/models"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	baseUrl     = "https://api.checks.truora.com/v1/checks"
	contentType = "application/x-www-form-urlencoded"
)

func TryToCreateWallet(userID int) (models.Log, error) {
	var err error

	user, err := US.GetUser(userID)
	if err != nil {
		return models.Log{}, err
	}

	canCreate, err := CheckIfCanCreateWallet(user)
	if err != nil {
		return models.Log{}, err
	}

	LogCreated, err := LS.CreateLog(&user, canCreate)
	if err != nil {
		return models.Log{}, err
	}

	return LogCreated, nil
}

// Function to handle the check_background consult at the Truora API
func CheckIfCanCreateWallet(user models.User) (bool, error) {

	checkID, err := postTruoraAPIRequest(&user)
	if err != nil {
		return false, err
	}

	time.Sleep(5 * time.Second)



	criminalRecordScore, err := getTruoraAPIRequest(checkID)
	if err != nil {
		return false, err
	}

	if criminalRecordScore < 1 {
		return false, nil
	}

	return true, nil
}

// Function to create a POST request in Truora API
func postTruoraAPIRequest(user *models.User) (string, error) {
	var err error
	var userInfo *strings.Reader

	if user.Country == "BR" {
		userInfo = strings.NewReader(fmt.Sprintf("national_id=%s&country=%s&type=person&date_of_birth=%s&user_authorized=true&force_creation=true", user.DocumentNumber, user.Country, user.DateOfBirth))
	} else {
		userInfo = strings.NewReader(fmt.Sprintf("national_id=%s&country=%s&type=person&user_authorized=true&force_creation=true", user.DocumentNumber, user.Country))
	}

	req, err := http.NewRequest("POST", baseUrl, userInfo)
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud para Truora API: %w", err)
	}

	req.Header.Add("truora-api-key", getAPI_KEY())
	req.Header.Add("content-type", contentType)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error enviar la solucitud a la API: %w", err)
	}

	defer res.Body.Close()

	var checkResponse models.TruoraPostResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer el contenido de la respuesta: %w", err)
	}

	err = json.Unmarshal(body, &checkResponse)
	if err != nil {
		return "", fmt.Errorf("error en decodificar la respuesta de la solicitud POST a la API: %w", err)
	}
	checkID := checkResponse.Check.CheckID

	return checkID, nil
}

// Function to create a GET request in Truora API
func getTruoraAPIRequest(checkID string) (int, error) {

	req, err := http.NewRequest("GET", baseUrl+checkID, strings.NewReader(""))
	if err != nil {
		return 0, fmt.Errorf("error al crear la solicitud para Truora API: %w", err)
	}

	req.Header.Add("truora-api-key", getAPI_KEY())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error enviar la solucitud a la API: %w", err)
	}

	defer res.Body.Close()

	var checkResult models.TruoraGetResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("error al leer el contenido de la respuesta: %w", err)
	}

	err = json.Unmarshal(body, &checkResult)
	if err != nil {
		return 0, fmt.Errorf("error al deserializar la respuesta JSON: %w", err)
	}

	criminalRecordScore, err := getScoreForCriminalRecords(&checkResult)
	if err != nil {
		return 0, err
	}

	return criminalRecordScore, nil
}

// Function to read the API_KEY in the .env file
func getAPI_KEY() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el archivo .env: %w", err)
	}
	apiKey := string(os.Getenv("TRUORA_API_KEY"))
	return apiKey
}

// Function to obtain the criminal records score
func getScoreForCriminalRecords(checkResult *models.TruoraGetResponse) (int, error) {
	for _, score := range checkResult.Check.Scores {
		if score.DataSet == "criminal_record" {
			return score.Score, nil
		}
	}
	return 0, fmt.Errorf("score del data-set 'criminal-records' no encontrado")
}
