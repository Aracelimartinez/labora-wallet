package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"labora-wallet/db"
	"labora-wallet/models"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	baseUrl     = "https://api.checks.truora.com/v1/checks"
	contentType = "application/x-www-form-urlencoded"
)

func TryToCreateWallet(user *models.User) (models.Log, error) {
	var err error

	canCreate, err := CheckIfCanCreateWallet(user)
	if err != nil {
		return models.Log{}, err
	}

	LogCreated, err := LS.CreateLog(user, canCreate)
	if err != nil {
		return models.Log{}, fmt.Errorf("no fue posible crear la solicitud: %w", err)
	}

	return LogCreated, nil
}

// Function to handle the check_background consult at the Truora API
func CheckIfCanCreateWallet(user *models.User) (bool, error) {

	checkID, err := postTruoraAPIRequest(user)
	if err != nil {
		return false, fmt.Errorf("error en la solicitud POST a la API truora: %w", err)
	}

	time.Sleep(5 * time.Second)

	criminalRecordScore, err := getTruoraAPIRequest(checkID)
	if err != nil {
		return false, fmt.Errorf("error en la solicitud GET a la API truora: %w", err)
	}

	if criminalRecordScore < 1 {
		return false, nil
	}

	return true, nil
}

// Function to create a POST request in Truora API
func postTruoraAPIRequest(user *models.User) (string, error) {
	var err error

	userInfo, err := setUserBodyRequest(user)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", baseUrl, userInfo)
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud POST para Truora API: %w", err)
	}

	req.Header.Add("truora-api-key", db.GlobalConfig.TruoraAPIKey)
	req.Header.Add("content-type", contentType)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", fmt.Errorf("error enviar la solucitud POST a la API: %w", err)
	}

	defer res.Body.Close()

	var checkResponse models.TruoraPostResponse
	var checkResponseError models.TruoraErrorResponse

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", fmt.Errorf("error al leer el contenido de la respuesta: %w", err)
	}

	if res.StatusCode != http.StatusCreated {

		err = json.Unmarshal(body, &checkResponseError)
		if err != nil {
			return "", fmt.Errorf("error en decodificar la respuesta de la solicitud: %w", err)
		}

		return "", fmt.Errorf(checkResponseError.Message)
	}

	err = json.Unmarshal(body, &checkResponse)
	if err != nil {
		return "", fmt.Errorf("error en decodificar la respuesta de la solicitud: %w", err)
	}

	checkID := checkResponse.Check.CheckID

	if checkID == "" {
		err = fmt.Errorf("checkID no vacio")
		return "", err
	}

	return checkID, nil
}

// Function to create a GET request in Truora API
func getTruoraAPIRequest(checkID string) (int, error) {

	req, err := http.NewRequest("GET", baseUrl +"/" + checkID, strings.NewReader(""))
	if err != nil {
		return 0, fmt.Errorf("error al crear la solicitud GET para Truora API: %w", err)
	}

	req.Header.Add("Truora-API-Key", db.GlobalConfig.TruoraAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, fmt.Errorf("error enviar la solucitud a la API: %w", err)
	}

	defer res.Body.Close()

	var checkResult models.TruoraGetResponse
	var checkResponseError models.TruoraErrorResponse

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("error al leer el contenido de la respuesta: %w", err)
	}

	if res.StatusCode != http.StatusOK {

		err = json.Unmarshal(body, &checkResponseError)
		if err != nil {
			return 0, fmt.Errorf("error en decodificar la respuesta de la solicitud: %w", err)
		}

		return 0, fmt.Errorf(checkResponseError.Message)
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

// Function to create the user info body request for the POST request to the Truora API
func setUserBodyRequest(user *models.User) (*strings.Reader, error) {
	var userInfo *strings.Reader
	if user.Country == "BR" {
		t, err := time.Parse("2006-01-02T15:04:05Z07:00", user.DateOfBirth)
		if err != nil {
			return nil, err
		}

		dateOfBirth := t.Format("02012006")

		log.Println(dateOfBirth)

		userInfo = strings.NewReader(fmt.Sprintf("national_id=%s&country=%s&type=person&date_of_birth=%s&user_authorized=true&force_creation=true", user.DocumentNumber, user.Country, dateOfBirth))

	} else {
		userInfo = strings.NewReader(fmt.Sprintf("national_id=%s&country=%s&type=person&user_authorized=true&force_creation=true", user.DocumentNumber, user.Country))

	}
	return userInfo, nil
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
