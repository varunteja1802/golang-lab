package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Database connection
const uri="mongodb://localhost:27017//?maxPoolSize=20&w=majority"

var client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
var coll = client.Database("db2").Collection("collection2")
func main(){
	
	getallmovies()
	getmovie()
	getmovie_time()
	getmovie_year()
	get_imdb()
	get_movie_not_USA()
	get_movie_USA_rated8()

}

// GET ALL MOVIES IN THE COLLECTION
func getallmovies() []bson.M{
	cur,err:= coll.Find(context.Background(),bson.D{{}})
	if err!=nil{
		log.Fatal(err)
	}

	var movies []bson.M

	for cur.Next(context.Background()){
		var movie bson.M
		err:=cur.Decode(&movie)
		if err!=nil{
			log.Fatal(err)
		}
		movies=append(movies, movie)
	}
	fmt.Println(movies)
	return movies
}

//GET THE MOVIE WITH PARTICULAR MOVIE NAME USING "$eq" 
func getmovie() []bson.M{
	filter := bson.M{"title": bson.M{"$eq":"A Corner in Wheat"}}
	curr,err:=coll.Find(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	var results []bson.M
	for curr.Next(context.Background()){
		var result bson.M
		err:= curr.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	fmt.Println(results)
	return results
}

//GET MOVIES WHICH ARE HAVING RUNTIME LESSTHAN 15 USING "$lte" 
func getmovie_time() []bson.M{
	filter:=bson.M{"runtime":bson.M{"$lte":15}}
	curr,err:=coll.Find(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	var results []bson.M
	for curr.Next(context.Background()){
		var result bson.M
		err:=curr.Decode(&result)
		if err!=nil{
			log.Fatal(err)
		}
		results=append(results, result)
	}
	fmt.Println(results)
	return results
}

//GET MOVIES WHICH ARE RELEASED AFTER 1917 USING "$gte"
func getmovie_year() []bson.M{
	filter:=bson.M{"year":bson.M{"$gte":1917}}
	curr,err:=coll.Find(context.Background(),filter)
	if err!=nil{
		log.Fatal(err)
	}
	var results []bson.M
	for curr.Next(context.Background()){
		var result bson.M
		err:=curr.Decode(&result)
		if err!=nil{
			log.Fatal(err)
		}
		results=append(results, result)
	}
	fmt.Println(results)
	return results
}

//Query on nested field; using dot notation 
func get_imdb() []bson.M{
	//Query on nested field; using dot notation 
	//("field.nestedField")
	filter:=bson.M{"imdb.rating":bson.M{"$gte":8}}
	curr,_:=coll.Find(context.Background(),filter)
	var results []bson.M
	for curr.Next(context.Background()){
		var result bson.M
		err:=curr.Decode(&result)
		if err!=nil{
			log.Fatal(err)
		}
		results=append(results, result)
	}
	fmt.Println(results)
	return results
}

//DE-SELECTING THE MOVIES WHICH ARE FROM USA 
func get_movie_not_USA() []bson.M{
	filter:=bson.M{"countries":bson.M{"$not":bson.M{"$eq":"USA"}}}
	curr,_:=coll.Find(context.Background(),filter)
	var results []bson.M
	for curr.Next(context.Background()){
		var result bson.M
		err:=curr.Decode(&result)
		if err!=nil{
			log.Fatal(err)
		}
		results=append(results, result)
	}
	fmt.Println(results)
	return results
}
}