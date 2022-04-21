package Summary_Classifier

import (
	"github.com/Darklabel91/Summary_Classifier"
	"github.com/Darklabel91/Summary_Classifier/Struct"
	"reflect"
	"testing"
)

func TestSummaryClassifier(t *testing.T) {
	type args struct {
		summary    string
		identifier string
		court      string
	}
	tests := []struct {
		name string
		args args
		want Struct.Infered_decision
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Summary_Classifier.SummaryClassifier(tt.args.summary, tt.args.identifier, tt.args.court); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SummaryClassifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryClassifierCSV(t *testing.T) {
	type args struct {
		rawDecisionPath  string
		separator        rune
		nameResultFolder string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Summary_Classifier.SummaryClassifierCSV(tt.args.rawDecisionPath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("SummaryClassifierCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
