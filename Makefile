build-populator:
	docker build db/populator -t db-populator:latest

run-populator:
	MONGOCONNSTRING="mongodb://db:27017/" docker run --rm -e MONGOCONNSTRING --network=projetoprogramacaoweb-back_private_network db-populator:latest