package httplib

import (
	"log"
	"net/http"
	"testing"
)

func TestRestFetcher_Get(t *testing.T) {
	respData := map[string]string{}

	restClient, err := NewRestClient("http://127.0.0.1:8888", "user")
	if err != nil {
		t.Fatal(err)
	}

	fetcher := restClient.GetFetcher("greet", nil)
	if fetcher.Get(map[string]interface{}{"name": "jimmy"}).Error() != nil {
		t.Fatal(fetcher.Err)
	}

	if statusCode, err := fetcher.Result(&respData); err != nil {
		t.Fatal(err)
	} else if statusCode != http.StatusOK {
		t.Fatalf("http status code  %d != %d", statusCode, http.StatusOK)
	}
	t.Log(respData)
}

func TestRestFetcher_Post(t *testing.T) {
	respData := map[string]string{}

	restClient, err := NewRestClient("http://127.0.0.1:8888", "")
	if err != nil {
		t.Fatal(err)
	}

	fetcher := restClient.GetFetcher("greet", nil)
	data := map[string]interface{}{"name": "bode"}
	if fetcher.Post(data).Error() != nil {
		t.Fatal(fetcher.Err)
	}

	if statusCode, err := fetcher.Result(&respData); err != nil {
		t.Fatal(err)
	} else if statusCode != http.StatusOK {
		t.Fatalf("http status code  %d != %d", statusCode, http.StatusOK)
	}
	t.Log(respData)
}

func BenchmarkRestFetcher_Get(b *testing.B) {

	for n := 0; n < b.N; n++ {
		respData := map[string]string{}
		restClient, err := NewRestClient("http://localhost:9000", "/")
		if err != nil {
			log.Println(err)
			b.Fatal(err)
		}
		fetcher := restClient.GetFetcher("greet", nil)
		if fetcher.Get(map[string]interface{}{"name": "jimmy"}).Error() != nil {
			log.Println(fetcher.Err)
			b.Fatal(fetcher.Err)
		}

		if statusCode, err := fetcher.Result(&respData); err != nil {
			log.Println(err)
			b.Fatal(err)
		} else if statusCode != http.StatusOK {
			b.Fatalf("http status code  %d != %d", statusCode, http.StatusOK)
		}
		//b.Log(respData)
	}

}
