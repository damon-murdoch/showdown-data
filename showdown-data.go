package main

// Operating System Library
import "os"

// Json read/write library
import "encoding/json"

// Input / Output Library
import "fmt"

// IO Utility Library
import "io/ioutil"

// String Library
import "strings"

// Return content from
// json object

func getJson(file *os.File)(map[string]interface{}){

	// Read the content from the file
	bytes, _ := ioutil.ReadAll(file);

	// Create a map for the result
	var result map[string]interface{};

	// Unmarshall the json content
	json.Unmarshal([]byte(bytes), &result);

	// Return the result to the terminal
	return result;
}

// Main Function
func main(){

	// Retrieve CLI Arguments
	args := os.Args[1:];

	if len(args) >= 1 {

		// Current Arguments:
		// ability, move, dex, type

		if args[0] == "ability" { // Ability Data
			
			// Open the ability json file
			file, err := os.Open("data/ability.json");

			// If the error is not null
			if err != nil {
				// Print it and exit 
				fmt.Println(err); 
			} else {
				// Get the Json data
				var data = getJson(file);

				// If we have more than two arguments
				if len(args) > 1 {

					// Make the search term lower case
					var search = strings.Replace(strings.ToLower(args[1]),"-","",-1);

					// Use the second argument as a search key
					var instance = data[search]

					// Dereference the instance 
					inst, ok := instance.(map[string]interface{});

					// If the dereference is ok
					if ok {

						// Print the Name
						fmt.Println("Name:",inst["name"]);

						// Print the Description
					}
				} else {
					fmt.Println("Usage: showdown-data ability [ability-name]");
				}
			}

			// Defer the closing so we can parse it later on
			defer file.Close();

		} else if args[0] == "move" { // Move Data

			// Open the move json file
			file, err := os.Open("data/move.json");

			// If the error is not null
			if err != nil {
				// Print it and exit 
				fmt.Println(err); 
			} else {
				// Get the Json data
				var data = getJson(file);

				// If we have more than two arguments
				if len(args) > 1 {

					// Make the search term lower case
					var search = strings.Replace(strings.ToLower(args[1]),"-","",-1);

					// Use the second argument as a search key
					var instance = data[search]

					// Dereference the instance 
					inst, ok := instance.(map[string]interface{});

					// If the dereference is ok
					if ok {

						// Print the name
						fmt.Println("Name:",inst["name"]);

						// Print the Accuracy
						fmt.Println("Accuracy:",inst["accuracy"]);

						// Print the Base Power
						fmt.Println("Base Power:",inst["basePower"]);

						// Print the Category
						fmt.Println("Category:",inst["category"]);

						// Print the Contest Type
						fmt.Println("Contest Type:",inst["contestType"]);

						// Print the Power Points
						fmt.Println("Power Points:",inst["pp"]);

						// Print the Target
						fmt.Println("Target:",inst["target"]);

						// Print the Type
						fmt.Println("Type:",inst["type"]);
						
						// Print the Priority
						fmt.Println("Priority:",inst["priority"]);

						// Print the Flags
						fmt.Println("Flags:",inst["flags"]);

						// Print the Description
					}
				} else {
					fmt.Println("Usage: showdown-data move [move-name]");
				}
			}

			// Defer the closing so we can parse it later on
			defer file.Close();

		} else if args[0] == "dex" { // Pokedex Data

			// Open the pokedex json file
			file, err := os.Open("data/dex.json");

			// If the error is not null
			if err != nil {
				// Print it and exit 
				fmt.Println(err); 
			} else {
				// Get the Json data
				var data = getJson(file);

				// If we have more than two arguments
				if len(args) > 1 {

					// Make the search term lower case
					var search = strings.Replace(strings.ToLower(args[1]),"-","",-1);

					// Use the second argument as a search key
					var instance = data[search]

					// Dereference the instance 
					inst, ok := instance.(map[string]interface{});

					// If the dereference is ok
					if ok {

						// sfmt.Println(inst);

						// Print the Name
						fmt.Println("Name:",inst["name"]);

						// Print the Types
						fmt.Println("Types:",inst["types"]);

						// Print the Abilities
						fmt.Println("Abilities:",inst["abilities"]);

						// Print the Base Stats
						fmt.Println("Base Stats:",inst["baseStats"]);

						// Print the Egg Groups
						fmt.Println("Egg Groups:",inst["eggGroups"]);

						// Print the Evolutions
						fmt.Println("Evolutions:",inst["evos"]);

						// Print the Other Formes
						fmt.Println("Other Formes:",inst["otherFormes"]);

						// Print the Gender Ratio
						fmt.Println("Gender Ratio:",inst["genderRatio"]);

						// Print the Height
						fmt.Println("Height:",inst["heightm"],"Meters");

						// Print the Weight
						fmt.Println("Weight:",inst["weightkg"],"Kilograms");

						// Print the Color
						fmt.Println("Colour:",inst["color"]);

						// Print the Description
					}
				} else {
					fmt.Println("Usage: showdown-data dex [pokemon-name]");
				}
			}

			// Defer the closing so we can parse it later on
			defer file.Close();

		} else if args[0] == "type" { // Type Data
			// Open the type json file
			file, err := os.Open("data/type.json");

			// If the error is not null
			if err != nil {
				// Print it and exit 
				fmt.Println(err); 
			} else {
				// Get the Json data
				var data = getJson(file);

				// If we have more than two arguments
				if len(args) > 1 {

					var search = strings.Title(strings.ToLower(args[1]));

					// Use the second argument as a search key
					var instance = data[search];
					
					// Dereference the instance 
					inst, ok := instance.(map[string]interface{});

					// If the dereference is ok
					if ok {

						// fmt.Println(instance);

						// Print the Name
						fmt.Println("Name:",search);

						// Optimal Hidden Power IVs
						fmt.Println("Hidden Power IVs:",inst["HPivs"]);

						// Optinal Hidden Power DVs
						fmt.Println("Hidden Power DVs:",inst["HPdvs"]);

						// Damage Taken
						fmt.Println("Damage Taken",inst["damageTaken"]);

						// 0: Neutral 1: Weak 2: Resist 3: Immune
						fmt.Println("0: Neutral 1: Weak 2: Resist 3: Immune");

						// Print the Description
					}
				} else {
					fmt.Println("Usage: showdown-data type [type-name]");
				}
			}

			// Defer the closing so we can parse it later on
			defer file.Close();
		}	
	} else {
		fmt.Println("Usage: showdown-data [category] [item]");
	}
}
