package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

const TRUCK_NUM = 3

type Box struct {
	Id     int
	Weight int
}

func main() {
	flag.Parse()

	boxes, _ := ParseBoxes(flag.Args())

	var (
		T [TRUCK_NUM][]Box
	)

	for _, box := range boxes {
		AssignToTruck(&T, box)
	}

	for i, truck := range T {
		fmt.Printf("truck_%d:%s\n", i+1, ShippingIds(truck))
	}
}

func ParseBoxes(flags []string) ([]Box, error) {
	var boxes []Box
	for _, boxVal := range flags {
		box := strings.Split(boxVal, ":")
		boxId, _ := strconv.Atoi(box[0])
		boxWeight, _ := strconv.Atoi(box[1])
		boxes = append(boxes, Box{
			Id:     boxId,
			Weight: boxWeight,
		})
	}
	return boxes, nil
}

// 箱をトラックに割り当てる
func AssignToTruck(trucks *[TRUCK_NUM][]Box, box Box) {

	for i := 0; i < TRUCK_NUM; i++ {
		boxes := trucks[i]

		// まだ箱を積んでいないトラックに割り当てる
		if len(boxes) == 0 {
			var new []Box
			new = append(new, box)
			trucks[i] = new
			return
		}

		// まだ箱を積んでいないトラックがある場合
		if isVacant(*trucks) {
			continue
		}

		sIndex := lightestTruckNum(*trucks)
		trucks[sIndex] = append(trucks[sIndex], box)
		return
	}
}

// 積んでいる箱の重さが一番軽いトラック番号を特定する
func lightestTruckNum(trucks [TRUCK_NUM][]Box) int {

	var (
		sIndex int
		sVal   int = -1
	)

	for i := 0; i < TRUCK_NUM; i++ {
		truckWeght := calcTruckWeight(trucks[i])

		if sVal == -1 {
			sVal = truckWeght
			sIndex = i
			continue
		}

		if sVal > truckWeght {
			sVal = truckWeght
			sIndex = i
		}
	}
	return sIndex
}

// トラック1台に積んでいる箱の重さを計算
func calcTruckWeight(boxes []Box) int {

	var totalWeight int
	for _, box := range boxes {
		totalWeight += box.Weight
	}
	return totalWeight
}

// 箱を積んでいないトラックの確認
func isVacant(trucks [TRUCK_NUM][]Box) bool {
	for _, boxes := range trucks {
		if len(boxes) == 0 {
			return true
		}
	}
	return false
}

// カンマ区切りの出荷ID
func ShippingIds(boxes []Box) string {
	var ids []string
	for _, box := range boxes {
		ids = append(ids, strconv.Itoa(box.Id))
	}
	return strings.Join(ids, ",")
}
