## Hyperledger Fabric Network Setup Documentation

### Step 1: Generate Artifacts

Navigate to the 'artifacts' folder and access the 'channel' subdirectory:

> cd artifacts/channel

Execute the script to generate crucial artifacts for the blockchain network:

> ./create-artifacts.sh

This script orchestrates the creation of essential components, including:

- Genesis Block: The initial block in the blockchain, establishing the network's foundation.
- Channel Transaction: The transaction detailing the creation of the channel, ensuring a secure communication path.
- MSP Configurations: Membership Service Provider configurations for both Org1 and Org2, defining the rules and policies for network participants.
- Certificates: Certificates for peers and orderers, facilitating secure communication within the network.

### Step 2: Run Docker Services

Move to the parent folder and initiate the Docker services:

> cd ..
> docker-compose up -d

This command activates various services within Docker containers, encompassing Certificate Authority (CA) services, peer services for both organizations, and ordering services employing the Raft consensus mechanism.

### Step 3: Create and Join Channel

Move one level down and execute the 'createChannel.sh' script:

> ./createChannel.sh

This script performs the following tasks:

- Creates the channel, establishing a communication pathway for network participants.
- Joins peers to the channel, enabling them to engage in transactions.
- Updates anchor peers, ensuring seamless communication by storing information about other peers in the network.

### Step 4: Deploy Chaincode

With the channel in place, deploy the chaincode by running:

> ./deploychaincode.sh

### Step 5: Generate Connection Profile

Move inside `artifacts/channel/` and run:

> ./ccp-generate.sh

This will generate the connection profile for both organizations.

### Step 6: Start API Server

Navigate to the `api-server` folder and run:

> npm install
> npm start

This will start the API server.


Note:
Locate the chaincode in the "artifacts/src/github.com/fabcar/javascript/lib" directory.

packageChaincode
sleep 2
installChaincode
sleep 2
queryInstalled1
sleep 2
approveForMyOrg1
sleep 2
queryInstalled2
sleep 2
approveForMyOrg2
sleep 2
commitChaincodeDefination

npm install -g http-server
http-server


couchDB - http://localhost:5984/_utils/
