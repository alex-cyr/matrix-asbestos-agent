# Matrix Asbestos Agent: Multi-Agentic AI Framework for Hazardous Materials

**Developer:** alex-cyr + e-alya
**Domain Partnership:** Built in collaboration with Matrix Engineering Group Certified Asbestos Inspectors.

## Overview
The Matrix Asbestos Agent is a proprietary A2A (Agent-to-Agent) microservice pipeline engineered to autonomously parse PLM/TEM Laboratory Analytical Reports and compile NESHAP/AHERA compliant Asbestos Demolition/Renovation Surveys.

Built on the same enterprise-grade Golang + Vertex AI architecture as the Matrix ESA Agent, this system pivots the logic engine from environmental geophysics to hazardous material regulation.

## The Asbestos Pipeline (A2A Network)
1. **Lab Parser Agent:** Ingests raw PLM lab reports and Chain of Custody (COC) PDFs. Extracts Sample IDs, Material Descriptions, and strict Asbestos % classifications (Chrysotile, Amosite, etc.).
2. **Material Assessor Agent:** Groups samples into Homogeneous Areas (HAs) and categorizes physical state (Friable vs. Non-Friable Category I/II).
3. **NESHAP Synthesizer Agent:** The legal engine. Classifies >1% asbestos findings as ACM/RACM and drafts the mandatory state/federal abatement recommendations.
4. **Template Compiler Agent:** Injects the verified JSON payload directly into the XML of Matrix's static Asbestos Survey Word Template.

## CEO / Domain Expert TODO List (e-alya):
1. **The Template:** Upload the `Matrix_Asbestos_Blank_Template.docx` into the `/knowledge` folder.
2. **The Variables:** Create a list of all the `{{Bracket_Tags}}` needed for the JSON compiler.
3. **The Logic:** Write the `.md` files in `.agents/skills/` to teach the AI the Matrix rules for Asbestos (e.g., when to recommend abatement vs. an O&M plan, the >1% rule).
4. **The Data:** Drop raw Lab Report PDFs into the input folder for initial `main` branch Antigravity testing.
