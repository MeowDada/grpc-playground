package main

import (
	"log"
	
	pb "github.com/meowdada/grpc-playground/go/proto"
	"github.com/golang/protobuf/proto"
)

func main() {
	
	army := &pb.Army{
		Warriors: []*pb.Warrior {
			&pb.Warrior {
				Name: "Jack",
				Weapons: []*pb.Warrior_Weapon {
					&pb.Warrior_Weapon{
						Name: "Sword",
						Type: pb.Warrior_SWORD,
						Damage: 70,
					},
					&pb.Warrior_Weapon{
						Name: "Knife",
						Type: pb.Warrior_KNIFE,
						Damage: 45,
					},
				},
			},
			&pb.Warrior {
				Name: "Alen",
				Weapons: []*pb.Warrior_Weapon {
					&pb.Warrior_Weapon{
						Name: "Knife",
						Type: pb.Warrior_KNIFE,
						Damage: 25,
					},
				},
			},
		},
	}

	log.Println(army)

	out, err := proto.Marshal(army)
	if err != nil {
		log.Fatalln("Failed to encode army:", err)
	}

	armyDup := &pb.Army{}

	if err := proto.Unmarshal(out, armyDup); err != nil {
		log.Fatalln("Failed to decode army:", err)
	}

	log.Println(armyDup)
}