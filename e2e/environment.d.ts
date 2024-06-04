declare global {
  namespace NodeJS {
    interface ProcessEnv {
      RPC_URL: string;
      MNEMONIC_ALICE: string;
      ADDRESS_ALICE: string;
      MNEMONIC_BOB: string;
      ADDRESS_BOB: string;
    }
  }
}

export {};
