package general

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"

	"github.com/montanaflynn/stats"
)

func Get_summation(start_point int, end_point int, rule func(i int) float64) float64 {
	var rst float64
	for i := start_point; i <= end_point; i++ {
		rst += rule(i)
	}
	return rst
}

func Get_mean(data_set []float64) float64 {
	summation := Get_summation(0, len(data_set)-1,
		func(i int) float64 {
			return data_set[i]
		})
	return summation / float64(len(data_set))
}

func IsEven(num int) bool {
	return num%2 == 0
}

// It returns the median value, and first and second index of the median value
func Get_median(ordered_set []float64) (float64, int, int) {
	var first_median_index int
	var second_median_index int
	if IsEven(len(ordered_set)) {
		first_median_index = len(ordered_set)/2 - 1
		second_median_index = first_median_index + 1
	} else {
		first_median_index = (len(ordered_set) - 1) / 2
		second_median_index = first_median_index
	}
	return (ordered_set[first_median_index] + ordered_set[second_median_index]) / 2, first_median_index, second_median_index
}

func get_expression_for_standard_deviation(data_set []float64, is_data_sample bool) string {
	var result string
	var mean = Get_mean(data_set)
	result += "("
	result += "("
	for _, v := range data_set {
		result += fmt.Sprintf("(%f - %f)^2 +", v, mean)
	}
	result += "0"

	var denominator int
	if is_data_sample {
		denominator = len(data_set) - 1
	} else {
		denominator = len(data_set)
	}

	result += fmt.Sprintf(") / %d )^(1/2)", denominator)
	return result
}

func calculate_standard_deviation(data_set []float64, is_data_sample bool) float64 {
	mean := Get_mean(data_set)
	summation := Get_summation(0, len(data_set)-1, func(i int) float64 {
		return (mean - data_set[i]) * (mean - data_set[i])
	})

	var denominator float64
	if is_data_sample {
		denominator = float64(len(data_set) - 1)
	} else {
		denominator = float64(len(data_set))
	}
	return math.Sqrt(summation / denominator)
}

func frequency_table_to_value_table(frequency_table []map[string]float64) []float64 {
	value_table := []float64{}
	for _, val := range frequency_table {
		for i := 0; i < int(val["frequency"]); i++ {
			value_table = append(value_table, val["value"])
		}
	}
	return value_table
}

func sortFloat64s(data_set []float64) []float64 {
	copied := make([]float64, len(data_set))
	copy(copied, data_set)
	sort.Float64s(copied)
	return copied
}

func get_first_quartile(ordered_set []float64) float64 {
	_, _, second_index_of_q2 := Get_median(ordered_set)
	q1, _, _ := Get_median(ordered_set[:second_index_of_q2])
	fmt.Println("Q1 : ", q1)
	return q1
}

func get_third_quartile(ordered_set []float64) float64 {
	_, first_index_of_q2, _ := Get_median(ordered_set)
	q3, _, _ := Get_median(ordered_set[first_index_of_q2+1:])
	fmt.Println("Q3 : ", q3)
	return q3
}

// getting Interquartile Range of the data set
func get_IQR(ordered_set []float64) float64 {
	return get_third_quartile(ordered_set) - get_first_quartile(ordered_set)
}

func get_outliers(ordered_set []float64) []float64 {
	outliers := []float64{}
	q1 := get_first_quartile(ordered_set)
	q3 := get_third_quartile(ordered_set)
	iqr := q3 - q1

	minimum_threshold := q1 - 1.5*iqr
	maximum_threshold := q3 + 1.5*iqr
	for _, val := range ordered_set {
		if len(outliers) > 0 && outliers[len(outliers)-1] == val {
			continue
		}
		if val < minimum_threshold || val > maximum_threshold {
			outliers = append(outliers, val)
		}
	}

	return outliers
}

func temptemp(mean float64, standard_deviation float64) func(float64) float64 {
	return func(x float64) float64 {
		return stats.NormCdf(x, mean, standard_deviation)
	}
}

func select_random_element(slice []interface{}) interface{} {
	return slice[rand.Intn(len(slice))]
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
