package recipe

import (
	"fmt"
	"slices"
)

type empty struct{}

type Recipe struct {
	Title       string
	Author      string
	Ingredients []Ingredient
	Directions  []string
	Tags        map[string]empty
}

func NewRecipe(title string, author string) Recipe {
	return Recipe{Title: title, Author: author, Ingredients: []Ingredient{}, Directions: []string{}, Tags: make(map[string]empty)}
}

func MakeRecipe(title string, author string, Ingredients []Ingredient, Directions []string, tags map[string]empty) Recipe {
	return Recipe{
		Title:       title,
		Author:      author,
		Ingredients: Ingredients,
		Directions:  Directions,
		Tags:        tags}
}

func (this *Recipe) AddIngredient(numerator int, denominator int, unit string, name string) error {
	ingredient, err := MakeIngredient(numerator, denominator, unit, name)

	if err != nil {
		return fmt.Errorf("AddIngredient: failed to add ingredient: %w", err)
	}

	this.Ingredients = append(this.Ingredients, ingredient)

	return err
}

func (this *Recipe) RemoveIngredient(index int) error {
	if index < 0 || index > len(this.Ingredients)-1 {
		return fmt.Errorf("RemoveIngredient: invalid index number")
	}

	this.Ingredients = append(this.Ingredients[:index], this.Ingredients[index+1:]...)
	return nil
}

func (this *Recipe) SwapIngredients(indexArray []int) error {
	if len(indexArray) != len(this.Ingredients) {
		return fmt.Errorf("SwapIngredients: array lengths don't match")
	}

	for i := 0; i < len(indexArray); i++ {
		if !slices.Contains(indexArray, i) {
			return fmt.Errorf("SwapIngredients: array doesn't contain all indices")
		}
	}

    temp := slices.Clone(this.Ingredients)

	for i := 0; i < len(indexArray); i++ {
        this.Ingredients[i] = temp[indexArray[i]]
	}

	return nil
}

func (this *Recipe) AddDirection(direction string) {
	this.Directions = append(this.Directions, direction)
}

func (this *Recipe) RemoveDirection(index int) error {
	if index < 0 || index > len(this.Directions)-1 {
		return fmt.Errorf("RemoveDirection: invalid index number")
	}

	this.Directions = append(this.Directions[:index], this.Directions[index+1:]...)
	return nil
}

func (this *Recipe) SwapDirections(indexArray []int) error {
	if len(indexArray) != len(this.Directions) {
		return fmt.Errorf("SwapDirections: array lengths don't match")
	}

	for i := 0; i < len(indexArray); i++ {
		if !slices.Contains(indexArray, i) {
			return fmt.Errorf("SwapDirections: array doesn't contain all indices")
		}
	}

    temp := slices.Clone(this.Directions)

	for i := 0; i < len(indexArray); i++ {
        this.Directions[i] = temp[indexArray[i]]
	}

	return nil
}

func (this *Recipe) AddTag(tag string) {
    this.Tags[tag] = empty{}
}

func (this *Recipe) RemoveTag(tag string) {
	delete(this.Tags, tag)
}

func (this *Recipe) ScaleUpRecipe(scalar int) {
	for i, ingredient := range this.Ingredients {
		ingredient.Multiply(scalar)
		this.Ingredients[i] = ingredient
	}
}

func (this *Recipe) ScaleDownRecipe(scalar int) error {
	for i, ingredient := range this.Ingredients {
		var err error = ingredient.Divide(scalar)

		if err != nil {
			return fmt.Errorf("ScaleDownRecipe: failed to scale down recipe: %w", err)
		} else {
			this.Ingredients[i] = ingredient
		}
	}

	return nil
}

func (this Recipe) String() string {
	var result string

	result += this.Title + "\n"
	result += "By: " + this.Author + "\n"
	result += "\n"

	if len(this.Tags) != 0 {
		result += "Tags: "

		for tag := range this.Tags {
			result += fmt.Sprintf("\"%v\" ", tag)
		}

		result += "\n"
	}

	for _, ingredient := range this.Ingredients {
		result += fmt.Sprintf("-  %v \n", ingredient.String())
	}

	result += "\n"

	for i, direction := range this.Directions {
		result += fmt.Sprintf("%d. %v \n", i+1, direction)
	}

	return result
}
