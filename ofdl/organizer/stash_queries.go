package organizer

type StashLookup interface {
	Count() int
	Medias() []StashMedia
}

type StashPerformer struct {
	ID   string
	Name string
	URL  string
}

type PerformerLookup struct {
	FindPerformers struct {
		Count      int
		Performers []StashPerformer
	} `graphql:"findPerformers(performer_filter:{name:{value:$name,modifier:INCLUDES}})"`
}

type (
	PerformerCreateInput map[string]interface{}
	PerformerUpdateInput map[string]interface{}
	PerformerCreate      struct {
		PerformerCreate StashPerformer `graphql:"performerCreate(input: $performer)"`
	}
	PerformerUpdate struct {
		PerformerUpdate StashPerformer `graphql:"performerUpdate(input: $performer)"`
	}
)

type StashMedia struct {
	ID   string
	Path string
}

type ImageLookup struct {
	FindImages struct {
		Count  int
		Images []StashMedia
	} `graphql:"findImages(image_filter:{path:{value:$path,modifier:INCLUDES}})"`
}

func (l ImageLookup) Count() int {
	return l.FindImages.Count
}

func (l ImageLookup) Medias() []StashMedia {
	return l.FindImages.Images
}

type ImageUpdate struct {
	ImageUpdate struct {
		ID string
	} `graphql:"imageUpdate(input: {id: $id, performer_ids: [$performerId], studio_id: $studioId, organized: true, title: $title, date: $date})"`
}

type SceneLookup struct {
	FindScenes struct {
		Count  int
		Scenes []StashMedia
	} `graphql:"findScenes(scene_filter:{path:{value:$path,modifier:INCLUDES}})"`
}

func (l SceneLookup) Count() int {
	return l.FindScenes.Count
}

func (l SceneLookup) Medias() []StashMedia {
	return l.FindScenes.Scenes
}

type SceneUpdate struct {
	SceneUpdate struct {
		ID string
	} `graphql:"sceneUpdate(input: {id: $id, performer_ids: [$performerId], studio_id: $studioId, organized: true, title: $title, date: $date})"`
}
