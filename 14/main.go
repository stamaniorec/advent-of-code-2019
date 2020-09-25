package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	Name     string
	Quantity int
}

type Reaction struct {
	Ingredients    []Ingredient
	ResultQuantity int
}

func getMinimumOre(
	reaction Reaction,
	reactions map[string]Reaction,
	quantityMultiplier int,
	leftovers map[string]int) int {

	var oreSum int

	for _, ingredient := range reaction.Ingredients {
		quantity := ingredient.Quantity * quantityMultiplier

		if leftover, ok := leftovers[ingredient.Name]; ok && leftover >= quantity {
			leftovers[ingredient.Name] -= quantity
			continue
		}

		if ingredient.Name == "ORE" {
			oreSum += quantity
		}

		reactionForIngredient := reactions[ingredient.Name]

		if leftover, ok := leftovers[ingredient.Name]; ok && leftover > 0 {
			quantity -= leftover
			leftovers[ingredient.Name] -= leftover
		}

		timesToRun := int(math.Ceil(float64(quantity) / float64(reactionForIngredient.ResultQuantity)))
		producedQuantity := timesToRun * reactionForIngredient.ResultQuantity
		oreSum += getMinimumOre(reactionForIngredient, reactions, timesToRun, leftovers)

		leftover := producedQuantity - quantity
		if leftover > 0 {
			_, found := leftovers[ingredient.Name]
			if found {
				leftovers[ingredient.Name] += leftover
			} else {
				leftovers[ingredient.Name] = leftover
			}
		}
	}

	return oreSum
}

func parseIngredient(input string) Ingredient {
	input = strings.TrimSpace(input)

	tokens := strings.Split(input, " ")
	quantity, _ := strconv.Atoi(strings.TrimSpace(tokens[0]))
	name := strings.TrimSpace(tokens[1])

	return Ingredient{
		Name:     name,
		Quantity: quantity,
	}
}

func parseReaction(line string) (string, Reaction) {
	var reaction Reaction

	tokens := strings.Split(line, "=>")
	ingredientsToken := tokens[0]
	outputToken := tokens[1]

	output := parseIngredient(outputToken)
	reaction.ResultQuantity = output.Quantity

	var ingredients []Ingredient
	for _, ingredient := range strings.Split(ingredientsToken, ",") {
		ingredients = append(ingredients, parseIngredient(ingredient))
	}
	reaction.Ingredients = ingredients

	return output.Name, reaction
}

func readReactions() (map[string]Reaction, error) {
	reactions := make(map[string]Reaction)

	f, err := os.Open("input.txt")
	if err != nil {
		return reactions, fmt.Errorf("Can not open input file: %w", err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		outputName, reaction := parseReaction(sc.Text())
		reactions[outputName] = reaction
	}

	if err := sc.Err(); err != nil {
		return reactions, fmt.Errorf("Error while reading input file: %w", err)
	}

	return reactions, nil
}

func GetMinimumOre(reactions map[string]Reaction, fuel int) int {
	fuelReaction := reactions["FUEL"]
	quantityMultiplier := fuel
	leftovers := make(map[string]int)

	return getMinimumOre(fuelReaction, reactions, quantityMultiplier, leftovers)
}

func main() {
	reactions, err := readReactions()
	if err != nil {
		fmt.Println(err)
		return
	}

	oreForOneFuel := GetMinimumOre(reactions, 1)
	fmt.Println(oreForOneFuel)

	const availableOre = 1e12

	// At least this much fuel can be produced
	low := availableOre / oreForOneFuel

	// Leftover materials can bring some improvement; 5 is ballparking it
	high := 5 * low

	maxFuel := low

	for low <= high {
		mid := (low + high) / 2

		oreNeeded := GetMinimumOre(reactions, mid)
		if oreNeeded <= availableOre && mid > maxFuel {
			maxFuel = mid
		}

		if oreNeeded > availableOre {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println(maxFuel)
}
