package main
import {
	"fmt"
	"math"
	"math/rand"
	"time"
}
type Individual struct {
	x float64
	fitness float64
}
type Population []Individual
func main(){
	rand.Seed(time.Now().UnixNano())
	var(
		populationSize = 100
		eliteCount = 5
		mutationProb = 0.05
		generation = 100
		minRange = -1.0
		maxRange = 2.0
	)
	population := generatePopulation(populationSize, minRange, maxRange)
	assignFitness(&population)
	for i: = 0; i < generation; i++ {
		sortPopulation(&population)
		newPopulation := Population{}
		//Elitism
		for j := 0; j < eliteCount; j++ {
			newPopulation = append(newPopulation, population[j])
		}
		//Crossover
		for j := eliteCount; j < populationSize; j++ {
			parent1 := population[selectParent(population)]
			parent2 := population[selectParent(population)]
			child := crossover(parent1, parent2)
			newPopulation = append(newPopulation, child)
		}
		//Mutation
		for j := eliteCount; j < populationSize; j++ {
			if rand.Float64() < mutationProb {
				mutate(&newPopulation[j], minRange, maxRange)
			}
		}
		population = newPopulation
		assignFitness(&population)
	}
	sortPopulation(&population)
	fmt.Println("Best solution: ", population[0].x)
}
func generatePopulation(size int, minRange float64, maxRange float64) Population {
	population := Population{}
	for i := 0; i < size; i++ {
		x := rand.Float64()*(maxRange-minRange) + minRange
		population = append(population, Individual{x: x})
	}
	return population
}
func assignFitness(population *Population) {
	for i := range *population {
		x := (*population)[i].x
		fitness := x*math.Sin(10*math.Pi*x) + 1.0
		(*population)[i].fitness = fitness
	}
}
func sortPopulation(population *Population) {
	sort.Slice(*population, func(i, j int) bool {
		return (*population)[i].fitness > (*population)[j].fitness
	})
}
func selectParent(population *Population) Individual{
	tournamentSize := 10
	tournament := make(Population, tournamentSize)
	for i := 0; i < tournamentSize; i++ {
		index := rand.Intn(len(*population))
		tournament[i] = (*population)[index]
	}
	sortPopulation(&tournament)
	return tournament[0]
}
func crossover(parent1 Individual, parent2 Individual) Individual {
	alpha := rand.Float64()
	child := Individual{
		x: alpha*parent1.x + (1-alpha)*parent2.x,
	}
	return child
}
func mutate(individual *Individual, minRange float64, maxRange float64) {
	individual.x = individual.x + rand.Float64()(maxRange-minRange) + minRange
	individual.y = individual.y + rand.Float64()*(maxRange-minRange) + minRange
}
//Crossover function
func (i *Individual) Crossover(other *Individual, crossoverRate float64) {
    if rand.Float64() < crossoverRate {
        crossoverIndex := rand.Intn(len(i.Genes))
        for j := crossoverIndex; j < len(i.Genes); j++ {
            i.Genes[j], other.Genes[j] = other.Genes[j], i.Genes[j]
        }
    }
}
//Mutation function
func (i *Individual) Mutation(mutationRate float64) {
    for j := 0; j < len(i.Genes); j++ {
        if rand.Float64() < mutationRate {
            i.Genes[j] = rand.Intn(2)
        }
    }
}