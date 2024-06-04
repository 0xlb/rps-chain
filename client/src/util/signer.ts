import { DirectSecp256k1HdWallet, OfflineDirectSigner } from "@cosmjs/proto-signing"

export const getSignerFromMnemonic = async (mnemonic: string): Promise<OfflineDirectSigner> => {
    return DirectSecp256k1HdWallet.fromMnemonic(mnemonic, {
        prefix: "rps",
    })
}