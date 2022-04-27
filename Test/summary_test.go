package Test

import (
	"github.com/Darklabel91/Summary_Classifier/Summary"
	"github.com/Darklabel91/Summary_Classifier/SummaryCSV"
	"reflect"
	"testing"
)

func TestClassify(t *testing.T) {
	type args struct {
		summary string
	}
	tests := []struct {
		name    string
		args    args
		want    Summary.InferredDecision
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Summary.Classify(tt.args.summary)
			if (err != nil) != tt.wantErr {
				t.Errorf("Classify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Classify() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnSummaryClass(t *testing.T) {
	type args struct {
		summary string
		char    int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Summary.ReturnSummaryClass(tt.args.summary, tt.args.char)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReturnSummaryClass() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReturnSummaryClass() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummaryClassifierCSV(t *testing.T) {
	type args struct {
		rawFilePath      string
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
			if err := SummaryCSV.SummaryClassifierCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("SummaryClassifierCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
