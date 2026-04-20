package core

import (
	"context"
	"fmt"
	"log/slog"
	"time"
	"cloud.google.com/go/vertexai/genai"
)

// Pipeline enforces the SequentialAgent pattern.
// Multi-Agent Microservice Pipeline: Core orchestrator for multi-agentic document processing.
type Pipeline struct {
	Agents    []*Agent
	Memory    []*Artifact
	ProjectID string
	Location  string
	SkipHITL  bool // System Override: Bypass manual UI review loops (for fully automated dev/testing)
}

func NewPipeline(project, location string, skipHITL bool, agents ...*Agent) *Pipeline {
	return &Pipeline{
		Agents:    agents,
		Memory:    make([]*Artifact, 0),
		ProjectID: project,
		Location:  location,
		SkipHITL:  skipHITL,
	}
}

// Run executes the chain sequentially, pausing on unapproved artifacts.
func (p *Pipeline) Run(ctx context.Context, initialPayload string) (string, error) {
	slog.Info("[INFO] System initialized: Enterprise Asbestos Pipeline: Initializing SequentialAgent Pipeline")

	currentPayload := initialPayload
	for i, agent := range p.Agents {
		slog.Info("[INFO] Agent Node execution started", "name", agent.Cfg.Name, "sequence_step", i+1)

		if i == 0 {
			// Optimal Pacing: The Parser just shoved 13 full PDFs into the engine. We wait exactly 1 minute here
			// so the Free-Tier TPM (Tokens-Per-Minute) bucket resets, then we blaze through the rest of the nodes instantly.
			slog.Warn("[WARN] Rate limit mitigation activated; pausing sequence to respect API quotas.
			time.Sleep(60 * time.Second)
		}

		artifact, err := agent.Execute(ctx, genai.Text(currentPayload))
		if err != nil {
			return "", fmt.Errorf("SYSTEM FAULT: Pipeline execution failed at node %s: %w", agent.Cfg.Name, err)
		}

		p.Memory = append(p.Memory, artifact)

		// Hard enforcement of Human-in-the-loop (Liability Mitigation pattern).
		if !artifact.Approved {
			if p.SkipHITL {
				slog.Warn("/// HITL MUTED /// Auto-approving unverified entity artifact due to active SkipHITL override.", "artifact_id", artifact.ID)
				artifact.Approved = true
			} else {
				slog.Warn("ENFORCING HITL YIELD: State paused for Validation by Matrix Engineering.",
					"artifact_id", artifact.ID,
					"agent", agent.Cfg.Name)

				// In production integration, the Vahalo Frontend Dashboard consumes this event and prompts the UI.
				// After manual approval (inline doc comments), the webhook calls the pipeline back.
				return "", fmt.Errorf("HITL (Human-In-The-Loop) Validation Pause: Entity Validation Required by Licensed Environmental Professional / QA Reviewer [Artifact %s]", artifact.ID)
			}
		}

		// A2A transmission: Accumulate state so downstream nodes retain the full context of prior nodes and initial extraction.
		currentPayload = currentPayload + "\n\n=== [MATRIX NODE: " + agent.Cfg.Name + "] ===\n" + artifact.Content
		slog.Info("/// ARTIFACT VERIFIED /// A2A Network Transmission to next matrix node initiating.", "bytes_transferred", len(currentPayload))
	}

	slog.Info("/// MATRIX EXECUTION COMPLETE /// Final Payload Ready for Extraction.")
	if len(p.Memory) > 0 {
		return p.Memory[len(p.Memory)-1].Content, nil
	}
	return currentPayload, nil
}
