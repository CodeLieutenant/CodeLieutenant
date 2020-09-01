package services

import (
	"log"
	"regexp"
	"testing"
)

func TestSignUrl(t *testing.T) {
	key := []byte("Test Key")

	signedUrl, err := SignUrl("http://localhost:3000/subscribe?email=%s", key, "test@gmail.com")

	if err != nil {
		t.Fatalf("Error while signing url: %v", err)
	}

	log.Printf("%s\n", signedUrl)

	matched, err := regexp.MatchString("^http://localhost:3000/subscribe\\?email=test%40gmail\\.com&signature=[a-zA-Z0-9-_]+$", signedUrl)

	if err != nil {
		t.Fatalf("Error in regex: %v", err)
	}

	if !matched {
		t.Fatalf("Signature URL does not match the RegExp")
	}

}


func TestVerify(t *testing.T) {
	key := []byte("Test Key")

	signedUrl := "http://localhost:3000/subscribe?email=test%40gmail.com&signature=aHR0cDovL2xvY2FsaG9zdDozMDAwL3N1YnNjcmliZT9lbWFpbD10ZXN0QGdtYWlsLmNvbQB5azgGygndjdUFS67Qb51rYaQT6Jk3huTeOYBMIiLE"

	if Verify(signedUrl, key) {
		t.Fatalf("Invalid Signature\n")
	}
}