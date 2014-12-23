package ngsi10

import (
	"strings"
	"github.com/LeandroGuillen/cedalion/ngsi"
)

/* ****************************************************************************
*
* UpdateContextRequest - 
*/
type UpdateContextRequest struct {
	contextElementVector	ngsi.ContextElementVector    
	updateActionType 		ngsi.UpdateActionType     
	Nombre	string
}

func (ucr *UpdateContextRequest) Init(){
	
}
func (ucr *UpdateContextRequest) Render(requestType ngsi.RequestType, format ngsi.Format, indent string) string {
	s := []string{
		"this",
		"is",
		"a",
		"joined",
		"string\n"};
		
	return strings.Join(s, "\n")
}
func (ucr *UpdateContextRequest) Check(requestType ngsi.RequestType, format ngsi.Format, indent string, predetectedError string, counter int) string {
	return "OK"
}
func (ucr *UpdateContextRequest) Release() {
	ucr.contextElementVector.Release()
}
func (ucr *UpdateContextRequest) Present(indent string) {
	
}
