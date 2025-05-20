import { faker, fakerPT_PT } from '@faker-js/faker'
import { MongoClient } from 'mongodb'

const client = new MongoClient(process.env.MONGOCONNSTRING)

await client.connect()

const database = client.db("ProjetoInternet")

const usersCollection = database.collection("users")

const eventsCollection = database.collection("events")
if(await eventsCollection.countDocuments() == 0) {
    const eventsToAdd = []
    
    for(let i = 0; i < 5; i++) {
        eventsToAdd.push({
            name: fakerPT_PT.person.jobTitle(),

        })
    }
    console.log(eventsToAdd);
}

process.exit(0);