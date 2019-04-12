package data

import (
	"file/skate/models"
	"file/skate/sql"
	"log"
)

type OrganizeModel struct {
}

func NewOrganizeModel() *OrganizeModel {
	return new(OrganizeModel)
}

func (o *OrganizeModel) GetAllPlayerById(oid string) []*models.SPlayer {
	engine := sql.GetSqlEngine()
	data := models.MorePlayer()
	err := engine.Table("s_player").
		Where("organize_id=?", oid).
		Cols("id", "player_name", "gender", "group_type").
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
func (o *OrganizeModel) CheckOrganizeLogin(username string, password string) (*models.SOrganize, bool) {
	engine := sql.GetSqlEngine()
	organize := models.NewOrganize()
	flag, err := engine.Where("username=? and password=?", username, password).Get(organize)
	if err != nil {
		log.Print(err.Error())
	}
	return organize, flag
}

func (o *OrganizeModel) GetAllPlayerScore(oid string) []*models.OrganizePlayerScore {
	engine := sql.GetSqlEngine()
	data := models.MoreOrganizePlayerScore()
	err := engine.Table("s_organize").
		Join("INNER", "s_player", "s_organize.id=s_player.organize_id").
		Join("INNER", "s_score", "s_score.player_id=s_player.id").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Cols("player_name", "s_group", "match_id", "match_type", "date", "time_score", "match_name", "group_type", "no").
		Where("s_organize.id=?", oid).
		Asc("s_score.id").
		Find(&data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (o *OrganizeModel) GetBestMatchScore(oid string, matchName string) *models.OrganizePlayerScore {
	engine := sql.GetSqlEngine()
	data := models.NewOrganizePlayerScore()
	_, err := engine.Table("s_organize").
		Join("INNER", "s_player", "s_organize.id=s_player.organize_id").
		Join("INNER", "s_score", "s_score.player_id=s_player.id").
		Join("INNER", "s_match", "s_score.match_id=s_match.id").
		Cols("player_name", "s_group", "match_id", "match_type", "date", "time_score", "match_name", "s_match.group_type", "no").
		Where("s_organize.id=? and s_match.match_name=? and s_score.time_score <> ? and s_score.time_score <> ?", oid, matchName, "00:00.000", "完成比赛").
		Asc("s_score.time_score").
		Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}

func (o *OrganizeModel) GetOrganizeNameById(oid string) *models.SOrganize {
	engine := sql.GetSqlEngine()
	data := models.NewOrganize()
	_, err := engine.Table("s_organize").
		Where("id=?", oid).
		Cols("organize_name").
		Get(data)
	if err != nil {
		log.Print(err.Error())
	}
	return data
}
