package Test_pack;

import org.hyperledger.fabric.contract.Context;
import org.hyperledger.fabric.contract.ContractInterface;
import org.hyperledger.fabric.shim.ChaincodeException;
import org.hyperledger.fabric.shim.ChaincodeStub;
import com.owlike.genson.Genson;


public class Chaincode_test implements ContractInterface {
    private final Genson genson = new Genson();

    // 实现智能合约接口的invoke方法
    public void invoke(Context ctx) {
        ChaincodeStub stub = ctx.getStub();
        // 获取调用的函数名称
        String function = stub.getFunction();

        // 根据函数名称进行不同的处理逻辑
        switch (function) {
            case "storePoints":
                storePoints(ctx, stub.getParameters());
                break;
            case "getPoints":
                getPoints(ctx, stub.getParameters());
                break;
            default:
                // 不支持的函数名称，抛出异常
                throw new ChaincodeException("Unsupported function: " + function);
        }
    }

    // 存储积分到区块链状态中
    private void storePoints(Context ctx, String[] args) {
        ChaincodeStub stub = ctx.getStub();

        if (args.length != 2) {
            // 参数不正确，抛出异常
            throw new ChaincodeException("Incorrect number of arguments. Expecting 2.");
        }

        String userId = args[0];
        int points = Integer.parseInt(args[1]);

        // 将积分存储到区块链状态中
        stub.putStringState(userId, String.valueOf(points));

        System.out.println("Points stored successfully!");
    }

    // 从区块链状态中获取用户积分
    private void getPoints(Context ctx, String[] args) {
        ChaincodeStub stub = ctx.getStub();

        if (args.length != 1) {
            // 参数不正确，抛出异常
            throw new ChaincodeException("Incorrect number of arguments. Expecting 1.");
        }

        String userId = args[0];

        // 从区块链状态中获取用户积分
        String points = stub.getStringState(userId);

        System.out.println("Points for user " + userId + ": " + points);
    }

    // 主方法，用于启动智能合约
    public static void main(String[] args) {
        // 启动智能合约
        new Chaincode_test().start(args);
    }
}