import { faker } from '@faker-js/faker'
import { MongoClient } from 'mongodb'

const client = new MongoClient(process.env.MONGOCONNSTRING)

await client.connect()

const database = client.db("ProjetoInternet")

const usersCollection = database.collection("users")
if(await usersCollection.countDocuments() == 0) {
    console.log(await usersCollection.insertOne({
        Email: "teste@iscte-iul.pt",
        Password: "12345#"
    }));
}
class Event {
    constructor(Name, date, Description, Organizer, Tags, Image, Reviews) {
        if(typeof Name == String)
            this.Name = Name
        if(typeof date == Date)
            this.Date = date
        if(typeof Description == String)
            this.Description = Description
        if(typeof Organizer == String)
            this.Organizer = Organizer
        if(typeof Tags == typeof String[0])
            this.Tags = Tags //new Array<String>(0)
        if(typeof Image == String)
            this.Image = Image
        if(typeof Reviews == typeof {Rating: Number, Comment: String, SubmittedOn: Date})
            this.Reviews = Reviews
    }
}

const eventsCollection = database.collection("events")
if(await eventsCollection.countDocuments() == 0) {
    const eventsToAdd = []
    
    for(let i = 0; i < 5; i++) {
        const tags = []
        for(let j = 0; j < 3; j++) {
            tags.push(faker.commerce.productAdjective())
        }
        const reviews = []
        for(let j = 0; j < 10; j++) {
            reviews.push({
                Rating: faker.number.int({min: 1, max: 5}),
                Comment: faker.lorem.sentences(2),
                SubmittedOn: faker.date.between({ from: '2024-01-01', to: Date.now() })
            })
        }
        eventsToAdd.push({
            Name: faker.person.jobTitle(),
            Date: faker.date.future(),
            Description: faker.lorem.sentences(2),
            Organizer: faker.person.firstName(),
            Tags: tags,
            Reviews: reviews
        })
    }
    console.log(await eventsCollection.insertMany(eventsToAdd))
}

process.exit(0);