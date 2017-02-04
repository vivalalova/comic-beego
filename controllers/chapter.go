package controllers

type ChapterController struct {
	BaseController
}

func (this *ChapterController) Get() {

	catalogID := this.Ctx.Input.Param(":id")

	if catalogID != "" {
		this.find(catalogID)
	}
}

func (this *ChapterController) find(catalogID string) {

}
