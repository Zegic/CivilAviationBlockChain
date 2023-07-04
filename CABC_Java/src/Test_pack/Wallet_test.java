package Test_pack;

import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;

import org.hyperledger.fabric.gateway.Wallet;
import org.hyperledger.fabric.gateway.Wallets;

public class Wallet_test {
	public static void main(String[] args) {
		Path walletDirectory = Paths.get("wallet");
		try {
			Wallet wallet = Wallets.newFileSystemWallet(walletDirectory);
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}
}
