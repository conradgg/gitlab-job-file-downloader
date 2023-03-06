package webhooks

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cavaliergopher/grab/v3"
)

func Receiver() {
	http.HandleFunc("/"+os.Getenv("LISTEN_PATH"), handleWebhook)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("LISTEN_PORT"), nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	webhookData := make(map[string]interface{})
	err := json.NewDecoder(r.Body).Decode(&webhookData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("got webhook payload")
	if webhookData["build_status"] == "success" {
		a := fmt.Sprintf("%v", webhookData["build_id"])
		b, _ := strconv.ParseFloat(a, 64)
		var d int64 = int64(b)
		deployFile(d)
	}
}

func deployFile(job_id int64) {
	var url string = fmt.Sprintf("https://gitlab.com/api/v4/projects/%s/jobs/%d/artifacts/%s?private_token=%s", os.Getenv("PROJECT_ID"), job_id, os.Getenv("FOLDER_PATH"), os.Getenv("PROJECT_TOKEN"))
	downloadFile(os.Getenv("ENDPOINT"), url)
}

func downloadFile(filepath string, url string) {
	fmt.Println("Download started")
	resp, err := grab.Get(filepath, url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Download saved to", resp.Filename)
}
