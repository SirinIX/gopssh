package label

import (
	"fmt"
	"gopssh/pkg/cache"
	"strings"
)

var (
	errorInvalidLabel      = fmt.Errorf("invalid label")
	errorDuplicateLabelKey = fmt.Errorf("duplicate label key")
)

type LabelSelector struct {
	Adopt   map[string]string `json:"adopt" yaml:"adopt"`
	Discard map[string]string `json:"discard" yaml:"discard"`
}

func SelectInstances(labels string, instaces cache.Instances) (cache.Instances, error) {
	if labels == "" {
		return instaces, nil
	}

	labelSelector, err := NewLabelSelector(labels)
	if err != nil {
		return nil, err
	}
	
	return labelSelector.SelectInstances(instaces), nil
}

func NewLabelSelector(raw string) (*LabelSelector, error) {
	raw = strings.TrimSpace(raw)
	adopt := map[string]string{}
	discard := map[string]string{}

	labels := strings.Split(raw, ",")
	for _, l := range labels {
		if !strings.Contains(l, "=") {
			fmt.Printf("label %v is invalid, error: %v", l, errorInvalidLabel)
			return nil, errorInvalidLabel
		}
		// '=' √
		if strings.Contains(l, "!=") {
			dSplit := strings.Split(l, "!=")
			if len(dSplit) != 2 {
				fmt.Printf("label %v is invalid, error: %v", l, errorInvalidLabel)
				return nil, errorInvalidLabel
			}
			discard[dSplit[0]] = dSplit[1]
			continue
		}
		// '!=' ×, '=' √
		aSplit := strings.Split(l, "=")
		if len(aSplit) != 2 {
			fmt.Printf("label %v is invalid, error: %v", l, errorInvalidLabel)
			return nil, errorInvalidLabel
		}
		if _, ok := adopt[aSplit[0]]; ok {
			fmt.Printf("label key %v is duplicate, error: %v", aSplit[0], errorDuplicateLabelKey)
			return nil, errorDuplicateLabelKey
		}
		adopt[aSplit[0]] = aSplit[1]
	}

	return &LabelSelector{
		Adopt:   adopt,
		Discard: discard,
	}, nil
}

func (s *LabelSelector) SelectInstances(instances cache.Instances) cache.Instances {
	adopted := cache.Instances{}
	for _, inst := range instances {
		if inst.HasLabels(s.Adopt) {
			adopted = append(adopted, inst)
		}
	}

	if len(s.Discard) == 0 {
		return adopted
	}
	discarded := cache.Instances{}
	for _, inst := range adopted {
		if !inst.HasLabels(s.Discard) {
			discarded = append(discarded, inst)
		}
	}

	return discarded
}

func (s *LabelSelector) IsInstanceMatched(instance *cache.Instance) bool {
	if !instance.HasLabels(s.Adopt) {
		return false
	}
	if instance.HasLabels(s.Discard) {
		return false
	}

	return true
}
