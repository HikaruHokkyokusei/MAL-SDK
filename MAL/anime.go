package MyAnimeListSDK

import (
	"encoding/json"
	"fmt"
	"strings"

	MALModels "github.com/HikaruHokkyokusei/AnimeTree/MAL/Models"
)

var (
	animeFieldsTier1 = []string{
		"alternative_titles",
		"average_episode_duration",
		"created_at",
		"end_date",
		"genres",
		"id",
		"main_picture",
		"mean",
		"media_type",
		"my_list_status{finish_date,is_rewatching,num_episodes_watched,score,start_date,status,updated_at}",
		"start_date",
		"title",
	}
	animeFieldsTier2 = []string{
		"broadcast",
		"nsfw",
		"num_episodes",
		"num_favorites",
		"num_list_users",
		"num_scoring_users",
		"pictures",
		"popularity",
		"rank",
		"rating",
		"source",
		"start_season",
		"status",
		"studios",
		"synopsis",
		"updated_at",
	}
	animeFieldsTier3 = []string{
		"recommendations{%s}",
		"related_anime{%s}",
	}
)

func getAllFieldsForAnime() string {
	t1 := strings.Join(animeFieldsTier1, ",")
	t2 := strings.Join(animeFieldsTier2, ",")
	t3 := fmt.Sprintf(strings.Join(animeFieldsTier3, ","), t1, t1)
	return strings.Join([]string{t1, t2, t3}, ",")
}

func (c MyAnimeListClient) GetAnime(animeId int64) (*MALModels.Anime, error) {
	resp, statusCode, err := c.request(
		GET,
		fmt.Sprintf("/anime/%d", animeId),
		map[string]string{
			"fields": getAllFieldsForAnime(),
		},
		"",
	)
	if err != nil {
		return nil, err
	}

	if statusCode != 200 {
		errFormat := fmt.Sprint("Response status code: ", statusCode)
		if val, ok := (*resp)["error"]; ok {
			errFormat = fmt.Sprint(errFormat, ". Error: ", val)
		}
		if val, ok := (*resp)["message"]; ok {
			errFormat = fmt.Sprint(errFormat, ". Message: ", val)
		}
		return nil, fmt.Errorf(errFormat)
	}

	temp, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	anime := &MALModels.Anime{}
	return anime, json.Unmarshal(temp, anime)
}
