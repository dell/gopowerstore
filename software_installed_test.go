package gopowerstore

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	softwareInstalledMockURL = APIMockURL + apiSoftwareInstalledURL
)

var softwareInstalledID1 = "043294b6-b9b9-4adf-9e94-c227d24e8e4e"
var softwareInstalledID2 = "1add59ae-8302-4e2d-842f-337bd4c60e6c"

func TestClientIMPL_GetSoftwareInstalled(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	respData := fmt.Sprintf(`[{"id": "%s"}, {"id": "%s"}]`, softwareInstalledID1, softwareInstalledID2)
	httpmock.RegisterResponder("GET", softwareInstalledMockURL,
		httpmock.NewStringResponder(200, respData))
	softwareInstalled, err := C.GetSoftwareInstalled(context.Background())
	assert.Nil(t, err)
	assert.Len(t, softwareInstalled, 2)
	assert.Equal(t, softwareInstalledID1, softwareInstalled[0].ID)
}
