package gendoc_test

import (
	"github.com/nagytech/protokit/utils"

	"testing"

	"github.com/nagytech/protoc-gen-doc"
)

func BenchmarkParseCodeRequest(b *testing.B) {
	set, _ := utils.LoadDescriptorSet("fixtures", "fileset.pb")
	req := utils.CreateGenRequest(set, "Booking.proto", "Vehicle.proto")
	plugin := new(gendoc.Plugin)

	for i := 0; i < b.N; i++ {
		plugin.Generate(req)
	}
}
