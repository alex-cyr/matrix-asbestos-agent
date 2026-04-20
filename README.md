# Matrix Asbestos Agent: Multi-Agentic AI Framework for Hazardous Materials

**Developer:** alex-cyr + e-alya
**Status:** M&A Google Verified, NESHAP Domain Logic Injected
**Architecture:** 100% Golang 1337 Agent-to-Agent (A2A) Network

---

## 🏛️ System Architecture

This repository hosts an ultra-fast, multi-agent AI pipeline engineered to ingest raw analytical laboratory PDFs, parse Inspector field notes (Chain of Custody), and autonomously generate legally-defensible Asbestos Survey reports in `.docx` format within 45 seconds of physical lab completion.

The backend is written entirely in Golang, designed to be deployed as a stateless microservice on Google Cloud Run (`vertex-api-migration` branch), allowing real-time invocation from GCP Storage Buckets or Matrix Engineering's custom endpoints.

### Domain Mechanics: The 1% Rule Logic
The LLM structure is uniquely trained to combat specific regulatory liabilities:
*   **The Layer Trap:** Drywall materials frequently have multi-layered optical properties (A/B/C layers). The Lab Parser is explicitly constrained to evaluate *every* layer before concluding a sample is clean. Missing a single 2% layer in a C-stratum results in massive legal liability.
*   **The 1% Rule & EPA NESHAP:** The NESHAP Synthesizer strictly enforces EPA regulations. Findings `<1%` or `None Detected` are filtered. Findings `>1%` are strictly categorized by friability (Category I, Category II, or Friable RACM) to mathematically dictate required demolition/abatement procedures.

---

## 🤖 The 3-Agent Pipeline (A2A Loop)

1.  **Lab Parser Agent (`gemini-1.5-flash`)**
    *   **Input:** Eurofins/AES `.pdf` results + Chain of Custody `.pdf`.
    *   **Goal:** Identifies samples and stitches Lab quantitative data (`>1% Chrysotile`) directly to physical inspector locations (`Kitchen Drywall`) utilizing the Layer Trap logic constraint.

2.  **NESHAP Synthesizer Agent (`gemini-1.5-pro`)**
    *   **Input:** The raw array yielded by the Parser.
    *   **Goal:** The legal engine. Executes The 1% Rule math. Segregates clean materials from Asbestos Containing Materials (ACM). Identifies high-risk RACM elements versus non-friable construction materials. Calculates tables and graceful failures for fully-clean environments.

3.  **Template Compiler Agent (`gemini-1.5-flash`)**
    *   **Input:** The synthesized categorized array.
    *   **Goal:** Maps findings to dynamic `{{Bracket_Tags}}`. The Golang system then executes a direct XML `<w:p>` injection into `knowledge/Matrix_Asbestos_Blank_Template.docx`, circumventing unstable formatting issues.

---

## 📂 Repository File Structure 

*   `cmd/asbestosd/main.go` - The Antigravity Local Terminal execution instance (used for free, isolated debugging).
*   `cmd/api/main.go` - The stateless Google Cloud Run REST API service with file/bucket handlers (`vertex-api-migration` branch).
*   `.agents/skills/` - The strict, Markdown-coded prompt instructions maintaining the Domain Logic for each node.
*   `asbestos_input/` - The root folder where lab PDFs and static notes are dropped for `asbestosd` ingestion.
*   `knowledge/` - Houses the blank `.doc` or `.docx` Rockdale templates.
*   `output/` - Contains the yielded `.json` artifacts and the final `CLOUD_FINAL_REPORT.docx`.

## 🚀 Execution Parameters

**Local Execution (Antigravity Terminal):**
```powershell
go build -o tmp_asbestos.exe ./cmd/asbestosd
./tmp_asbestos.exe -payload "D:\path\to\your\matrix-asbestos-agent"
```

**Cloud Microservice Build Context:**
Built natively on Google Vertex AI REST integration (`google.golang.org/api/iterator`, `cloud.google.com/go/vertexai/genai`).
