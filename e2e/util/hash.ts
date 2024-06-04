import * as crypto from "crypto";

export const SALT = "my_test_salt"

export const generateHash = (move: string, salt: string): string => {
  const hash = crypto.createHash("sha256");
  hash.update(move + salt);
  return hash.digest("hex");
};