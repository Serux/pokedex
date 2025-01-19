package apiresponses

type Locationarea struct {
	Count    int
	Next     *string
	Previous *string
	Results  []struct {
		Name string
		Url  string
	}
}

type ExploreArea struct {
	Pokemon_encounters []struct {
		Pokemon struct {
			Name string
		}
	}
}

type Pokemon struct {
	Name   string
	Height int
	Weight int
	Stats  []struct {
		Base_Stat int
		Stat      struct {
			Name string
		}
	}
	Types []struct {
		Type struct {
			Name string
		}
	}
	Base_experience int
}
