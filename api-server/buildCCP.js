const { buildCCPOrg1 ,buildCCPOrg2} = require("./AppUtils");

exports.getCCP = (org) => {
    let ccp;
    switch (org) {
        case 1:
            ccp = buildCCPOrg1();
            break;
        case 2:
            ccp = buildCCPOrg2();
            break;
    }
    return ccp;
}