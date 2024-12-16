package main

import (
	"a21hc3NpZ25tZW50/service"
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// Initialize the services
var fileService = &service.FileService{}
var aiService = &service.AIService{Client: &http.Client{}}
var store = sessions.NewCookieStore([]byte("my-key"))

func getSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "chat-session")
	return session
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve the Hugging Face token from the environment variables
	token := os.Getenv("HUGGINGFACE_TOKEN")
	if token == "" {
		log.Fatal("HUGGINGFACE_TOKEN is not set in the .env file")
	}

	// Set up the router
	router := mux.NewRouter()

	// File upload endpoint
	router.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error reading file: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer file.Close()

		query := r.FormValue("question")
		if query == "" {
			http.Error(w, "Query is missing", http.StatusBadRequest)
			return
		}

		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			http.Error(w, "Error reading CSV: "+err.Error(), http.StatusInternalServerError)
			return
		}

		if len(records) == 0 {
			http.Error(w, "CSV file is empty", http.StatusBadRequest)
			return
		}

		table := make(map[string][]string)
		headers := records[0]
		for i := 1; i < len(records); i++ {
			for j, value := range records[i] {
				header := headers[j]
				table[header] = append(table[header], value)
			}
		}

		aiService := service.AIService{Client: http.DefaultClient}
		aiResponse, err := aiService.AnalyzeData(table, query, token)
		if err != nil {
			http.Error(w, "Error analyzing data with AI: "+err.Error(), http.StatusInternalServerError)
			return
		}

		respPayload := map[string]string{
			"answer": aiResponse,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respPayload)
	}).Methods("POST")

	// Chat endpoint
	router.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		// Ensure the request is a POST
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Parse the incoming JSON payload
		var payload struct {
			Query string `json:"query"`
		}

		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
			return
		}

		if payload.Query == "" {
			http.Error(w, "Query is required", http.StatusBadRequest)
			return
		}

		// Create an instance of AIService
		aiService := service.AIService{Client: http.DefaultClient}

		// Call ChatWithAI
		response, err := aiService.ChatWithAI("You are a helpful assistant.", payload.Query, token)
		if err != nil {
			http.Error(w, "Error querying AI: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Respond with the AI's answer
		respPayload := map[string]string{
			"answer": response.GeneratedText,
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(respPayload)
	}).Methods("POST")

	// Enable CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow your React app's origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	}).Handler(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler))
}
