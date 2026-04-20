package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"cloud.google.com/go/vertexai/genai"
	"github.com/matrix-engineering/matrix-esa-agent/internal/core"
)

// ExtractAsbestosPayload reads all PDFs in the suite directory and feeds them to the Parser Agent
func ExtractAsbestosPayload(ctx context.Context, pAgent *core.Agent, payloadDir string) (string, error) {
	log.Printf("[INFO] System initialized: Enterprise Asbestos Pipeline: Initiating Full Suite PDF Extraction for %s\n", payloadDir)

	// 1. Read all PDF files from the data folder
	entries, err := os.ReadDir(payloadDir)
	if err != nil {
		return "", fmt.Errorf("failed to read payload directory: %w", err)
	}

	var fullExtractedData string

	// 2. Define the strict extraction prompt based on our Logic Matrix
	extractionPrompt := `You are the PLM Lab Parser Agent. I have attached the Asbestos 3-File Payload (Eurofins Lab PDF, COC PDF, and Inspector Checklist).
You must execute high-fidelity data extraction:
1. Locate the Lab Results. Evaluate EVERY layer (Layer A, B, C) for asbestos %.
2. Cross-reference the Sample ID from the Lab Report with the physical location listed on the Chain of Custody (COC).
3. Extract the basic building metadata from the Checklist.
Return all extracted data as a structured JSON payload.`

	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".pdf" {
			continue
		}

		pdfPath := filepath.Join(payloadDir, entry.Name())
		log.Printf("/// INGESTING NODE PAYLOAD (SEQUENTIAL) ///: %s\n", entry.Name())

		pdfBytes, err := os.ReadFile(pdfPath)
		if err != nil {
			log.Printf("WARNING: failed to read EDR report %s: %v", entry.Name(), err)
			continue
		}

		var parts []genai.Part
		parts = append(parts, genai.Text(extractionPrompt))
		parts = append(parts, genai.Blob{MIMEType: "application/pdf", Data: pdfBytes})

		log.Printf("[INFO] Agent Node execution started: Parsing Document: %s\n", entry.Name())

		response, err := pAgent.Execute(ctx, parts...)
		if err != nil {
			log.Printf("WARNING: parser agent failed to process document %s: %v", entry.Name(), err)
			continue
		}

		fullExtractedData += "\n\n=== [DOCUMENT EXTRACT: " + entry.Name() + "] ===\n" + response.Content
	}

	if fullExtractedData == "" {
		return "", fmt.Errorf("no data extracted from any pdf files in payload directory")
	}

	log.Println("HITL (Human-In-The-Loop) Validation Pause: Sequential Suite Extraction Successful.")
	return fullExtractedData, nil
}
