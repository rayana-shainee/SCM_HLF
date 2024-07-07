const { Wallets } = require("fabric-network");
const FabricCAServices = require('fabric-ca-client');
const { buildCAClient, registerAndEnrollUser, enrollAdmin } = require("./CAUtil")
const { buildCCPOrg1,buildCCPOrg2,buildWallet } = require("./AppUtils");
const { getCCP } = require("./buildCCP");
const path=require('path');
const walletPath=path.join(__dirname,"wallet")
exports.registerUser = async ({ OrgMSP, userId }) => {

    let org = Number(OrgMSP.match(/\d/g).join(""));
    let ccp = getCCP(org)
    const caClient = buildCAClient(FabricCAServices, ccp, `ca.org${org}.example.com`);

    
    const wallet = await buildWallet(Wallets, walletPath);

   
    await enrollAdmin(caClient, wallet, OrgMSP);


    await registerAndEnrollUser(caClient, wallet, OrgMSP, userId, `org${org}.department1`);

    return {
        wallet
    }
}