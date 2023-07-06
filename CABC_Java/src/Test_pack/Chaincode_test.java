package Test_pack;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.shim.ChaincodeException;
import org.hyperledger.fabric.shim.ChaincodeStub;
import com.owlike.genson.Genson;
import java.util.List;
import org.everit.json.*;

public class Chaincode_test implements ContractInterface {
    private final Genson genson = new Genson();

    public void invoke(Context ctx) {
        ChaincodeStub stub = ctx.getStub();

        String function = stub.getFunction();
        switch (function) {
            case "storePoints":
                storePoints(ctx, stub.getParameters());
                break;
            case "getPoints":
                getPoints(ctx, stub.getParameters());
                break;
            default:
                throw new ChaincodeException("Unsupported function: " + function);
        }
    }

    private void storePoints(Context ctx, List<String> args) {
        ChaincodeStub stub = ctx.getStub();

        if (args.size() != 2) {
            throw new ChaincodeException("Incorrect number of arguments. Expecting 2.");
        }

        String userId = args.get(0);
        int points = Integer.parseInt(args.get(1));

        stub.putStringState(userId, String.valueOf(points));

        System.out.println("Points stored successfully!");
    }

    private void getPoints(Context ctx, List<String> args) {
        ChaincodeStub stub = ctx.getStub();

        if (args.size() != 1) {
            throw new ChaincodeException("Incorrect number of arguments. Expecting 1.");
        }

        String userId = args.get(0);

        String points = stub.getStringState(userId);

        System.out.println("Points for user " + userId + ": " + points);
    }

    public static void main(String[] args) {
        new Chaincode_test().invoke(new Context(null));
    }
}