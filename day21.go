package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Day21_1(filename string) int {
	fmt.Printf("")
	sc := NewFoodScanner(filename)
	sc.FindAllergens()
	fmt.Println(sc.String())
	return sc.CountIngredientsWithoutAllergen()
}

func Day21_2(filename string) int {
	mv := NewMsgValidator(filename)
	return mv.CountRule0()
}

type FoodScanner struct {
	ingredients map[string]*Ingredient
	allergens   map[string]*Allergen
}

func NewFoodScanner(filename string) FoodScanner {
	sc := new(FoodScanner)
	sc.ingredients = make(map[string]*Ingredient)
	sc.allergens = make(map[string]*Allergen)

	r_line := regexp.MustCompile(`([a-z ]*) \(contains ([a-z ,]*)\)`)
	for line := range inputCh(filename) {
		parsed := r_line.FindStringSubmatch(line)
		ingredients := strings.Split(parsed[1], " ")
		engs := strings.Split(parsed[2], ", ")
		sc.AddIngredients(ingredients)
		for _, eng := range engs {
			//fmt.Printf("%s -> %s\n", eng, ingredients)
			sc.AddAllergen(eng, ingredients)
		}
	}

	return *sc
}

func (sc *FoodScanner) AddIngredients(names []string) {
	for _, name := range names {
		sc.AddIngredient(name)
	}
}

func (sc *FoodScanner) AddIngredient(name string) int {
	if _, ok := sc.ingredients[name]; !ok {
		ingredient := Ingredient{name, 0, []string{}}
		sc.ingredients[ingredient.name] = &ingredient
	}
	sc.ingredients[name].count++
	return sc.ingredients[name].count
}

func (sc *FoodScanner) AddAllergen(name string, ingredients []string) {
	if allergen, ok := sc.allergens[name]; ok {
		if allergen.Combine(ingredients) == 1 {
			//sc.RemoveFromAllergens(allergen.ingredients[0])
		}
	} else {
		allergen := NewAllergen(name, ingredients)
		sc.allergens[allergen.name] = &allergen
	}
}

func (sc *FoodScanner) RemoveFromAllergens(ingredient string) bool {
	wasChange := true
	for _, allergen := range sc.allergens {
		if len(allergen.ingredients) == 1 {
			continue
		}
		allergen.ingredients = remove(allergen.ingredients, ingredient)
		wasChange = true
	}
	return wasChange
}

func (sc *FoodScanner) FindAllergens() {
	for {
		for _, allergen := range sc.allergens {
			ingredientName := allergen.GetIngredient()
			//fmt.Printf("'%s' -> '%s'\n", allergen.name, allergen.ingredients)
			if len(allergen.GetIngredient()) > 0 {
				sc.ingredients[ingredientName].allergens = append(sc.ingredients[ingredientName].allergens, allergen.name)
				sc.RemoveFromAllergens(ingredientName)
			}
		}
		if sc.allAllergensFound() {
			break
		}
	}
}

func (sc *FoodScanner) allAllergensFound() bool {
	for _, allergen := range sc.allergens {
		if len(allergen.ingredients) > 1 {
			return false
		}
	}
	return true
}

func (sc *FoodScanner) CountIngredientsWithoutAllergen() int {
	count := 0
	for _, ingredient := range sc.ingredients {
		if len(ingredient.allergens) == 0 {
			count += ingredient.count
		}
	}
	return count
}

func (sc *FoodScanner) String() string {
	s := "ingredients with allergen:\n"
	allergicIngredients := filterIngredients(sc.ingredients, func(ingredient *Ingredient) bool { return len(ingredient.allergens) > 0 })
	for _, ingredient := range allergicIngredients {
		s += fmt.Sprintf("   + %s: %s\n", ingredient.name, ingredient.allergens)
	}
	s += "ingredients without allergens:\n"
	for _, ingredient := range filterIngredients(sc.ingredients, func(ingredient *Ingredient) bool { return len(ingredient.allergens) == 0 }) {
		s += fmt.Sprintf("   + %s: %d\n", ingredient.name, ingredient.count)
	}
	return s
}

type Ingredient struct {
	name      string
	count     int
	allergens []string
}

type Allergen struct {
	name        string
	ingredients []string
}

func NewAllergen(name string, ingredients []string) Allergen {
	allergen := new(Allergen)
	allergen.name = name
	allergen.ingredients = ingredients
	return *allergen
}

func (allergen *Allergen) Combine(ingredients []string) int {
	inter := intersection(allergen.ingredients, ingredients)
	//fmt.Printf("    '%s' intersection of %s and %s -> %s\n", allergen.allergen, allergen.ingredients, ingredients, inter)
	//fmt.Printf("    '%s' intersection of %d and %d -> %d\n", allergen.name, len(allergen.ingredients), len(ingredients), len(inter))
	if len(inter) > 0 {
		allergen.ingredients = inter
	}
	return len(allergen.ingredients)
}

func (allergen *Allergen) GetIngredient() string {
	if len(allergen.ingredients) == 1 {
		return allergen.ingredients[0]
	}
	return ""
}

// TOOLS
func filterIngredients(ingredients map[string]*Ingredient, cond func(*Ingredient) bool) map[string]*Ingredient {
	filtered := make(map[string]*Ingredient, 0)
	for _, ingredient := range ingredients {
		if cond(ingredient) {
			filtered[ingredient.name] = ingredient
		}
	}
	return filtered
}
