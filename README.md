# Matrix Asbestos Compliance Engine: Multi-Agentic NESHAP Reporting

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Google Cloud](https://img.shields.io/badge/GCP-Vertex_AI-4285F4?style=flat&logo=googlecloud)](https://cloud.google.com/vertex-ai)
[![Architecture](https://img.shields.io/badge/Architecture-Agentic_AI-8A2BE2?style=flat)]()
[![Compliance](https://img.shields.io/badge/Compliance-EPA_NESHAP-008000?style=flat)]()

## Executive Summary: M&A Value and B2B SaaS Integration
The **Matrix Asbestos Compliance Engine** is an enterprise-grade, multi-agentic artificial intelligence system built on a Golang backend and integrated directly with Google's Vertex AI API. 

Designed specifically for the high-velocity demands of **Mergers & Acquisitions (M&A) due diligence** and commercial real estate transactions, this system eliminates the human administrative bottleneck inherent in environmental reporting. By autonomously translating raw laboratory data and field notes into legally defensible, perfectly formatted documentation, the architecture compresses a standard 24-to-48-hour engineering turnaround into a **45-second automated generation cycle**.

Engineered as a headless REST API utilizing Google Cloud Run, this system integrates seamlessly with cloud-based property management and reporting platforms like **Vahalo**. By routing the data pipeline directly into Vahalo's cloud infrastructure, field inspectors can upload Chain of Custody (COC) photos and project data from the job site, instantly triggering complex, NESHAP-compliant report generation via their mobile devices.

---

## 1. Core Enterprise Workflow & Liability Mitigation

To successfully automate hazardous materials reporting, the system relies on a deterministic, three-file payload to provide complete scientific and spatial context:
1. **The Laboratory Analytical Report (PDF):** Scientific data generated via Polarized Light Microscopy (PLM) detailing Sample IDs, material layers, and asbestos percentages.
2. **The Chain of Custody / COC (PDF):** The field inspector's notes, establishing the legal chain of possession and mapping the Lab Sample ID to its physical location.
3. **The Project Metadata Checklist (PDF):** Standardized metadata (site address, inspector details, project scope).

### 🚨 Liability Mitigation: "The Layer Trap" Protocol
The computational logic is rigorously trained to navigate stratified material compositions. Building materials (e.g., drywall systems) are tested in discrete layers (*Layer A: Joint Compound, Layer B: Paint, Layer C: Drywall*). Generic parsers create severe legal liability by only reading the top-level result. 

**Our system evaluates every single layer**. If any sub-layer tests >1% for asbestos, the entire Sample ID is mathematically flagged as Asbestos-Containing Material (ACM), neutralizing a severe liability blindspot common in manual data entry.

---

## 2. Spatial Data Structuring: AHERA & NESHAP Compliance

The spatial logic of the pipeline enforces the EPA's overarching **Homogeneous Area (HA)** contamination rules. If a single bulk sample within a designated Homogeneous Area tests positive for asbestos, the entire HA is legally classified as containing asbestos.

### The Regulatory Engine (NESHAP Synthesizer)
The AI utilizes the strict **EPA 1% Threshold Rule**:
* **Below Reportable Limits (Clean):** Lab results reading "None Detected" or `< 1%` are legally filtered out.
* **Actionable ACM:** Results strictly `> 1%` (e.g., 2% Chrysotile) are actionable. Actionable hits are automatically classified by friability (**Category I Non-Friable**, **Category II Non-Friable**, or high-risk **RACM**).
* **Volumetric Aggregation:** The AI aggregates estimated quantities (Square Feet / Linear Feet) to trigger state-level EPD project notifications prior to disturbance.

---

## 3. The 3-Agent Microservice Pipeline

The Agent-to-Agent (A2A) network utilizes a highly deterministic, sequential pipeline:

1. **Agent 1: The PLM Lab Parser (Gemini 1.5 Flash)**
   * *Objective:* High-Fidelity Data Extraction. Maps Sample IDs across Lab and COC documents, navigates multi-layered testing, and yields a unified JSON array.
2. **Agent 2: The NESHAP Synthesizer (Gemini 1.5 Pro)**
   * *Objective:* Regulatory Logic Engine. Consumes the Parser's array, applies the EPA 1% Rule, categorizes hits by Friability, and injects exclusionary boilerplate if a site is legally "clean."
3. **Agent 3: The Template Compiler (Gemini 1.5 Flash)**
   * *Objective:* Document Architecture. Maps synthesized findings exactly to `{{Bracket_Tags}}` required by the target `.docx` template.

---

## 4. Repository Topology & Deployment Strategy

This repository implements a unified Golang structure accommodating both local heuristic tuning and highly available cloud infrastructure.

### Local Development (`cmd/cli/main.go`)
Optimized for zero-cost prompt engineering and pipeline testing. Ingests local PDFs from a directory payload and monitors `stdout` for JSON parsing tolerances without incurring API server costs. 
```bash
go run cmd/cli/main.go -payload ./data/project_folder -skip-hitl=true
```

### Enterprise Cloud Run API (`cmd/api/main.go`)
Adapts the core architecture into a stateless, highly available API service containerized for Google Cloud Run. Accepts asynchronous POST requests containing document URIs from Google Cloud Storage, enabling direct webhook integration with platforms like Vahalo.

```json
// POST /api/v1/analyze/bucket
{
  "input_bucket": "matrix-vahalo-production-storage",
  "folder_prefix": "projects/asbestos/PRJ-99882/"
}
```

### Dynamic OOXML Document Compilation
The Go backend bypasses fragile word processing libraries by interacting directly with the `.docx` archive architecture. Using complex XML node cloning, the system dynamically scales and injects table rows (`<w:tr>`) to match the exact number of Homogeneous Areas found, preserving native document styles, shading, and borders.

### 🔒 Human-in-the-Loop (HITL) Safety Constraints
To mitigate legal liability, the system is engineered with a strict HITL Validation Checkpoint. Each AI-generated artifact initializes with an `Approved: false` struct state. In an enterprise deployment, this emits a webhook suspending execution until a Licensed Environmental Professional approves the array via the frontend interface. (Can be overridden via `-skip-hitl` flags for automated testing).

> Developed by Matrix Engineering Group. Proprietary framework engineered for high-fidelity compliance and due diligence reporting.
