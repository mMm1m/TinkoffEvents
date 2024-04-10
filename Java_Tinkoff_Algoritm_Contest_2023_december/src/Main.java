import java.util.*;

class Node {
    int id;
    int density;
    List<Node> neighbors;

    Node(int id, int density) {
        this.id = id;
        this.density = density;
        this.neighbors = new ArrayList<>();
    }
}

public class Main {

    // task1
    public static boolean isAnagram(final String s, String t) {
        char[] char_s = s.toCharArray();
        char[] char_t = t.toCharArray();
        Arrays.sort(char_s);
        Arrays.sort(char_t);
        if(s.length() != t.length()) return false;
        for(int i = 0; i < s.length(); ++i)
            if(char_s[i] != char_t[i]) return false;
        return true;
    }
    // task2

    public static boolean task2(String[] arr) {
        int[] res = Arrays.stream(arr).mapToInt(Integer::parseInt).toArray();
        List<Node> nodes = new ArrayList<>();
        for(int i = 0; i < res.length; ++i)
            nodes.add(new Node(i , res[i]));

        for (Node node : nodes) {
            List<Integer> neighbors = new ArrayList<>();
            for (int i = 0; i < node.density; i++) {
                if (i < nodes.size() && i != node.id) {
                    neighbors.add(i);
                }
            }
            if (neighbors.size() > node.density) {
                return false;
            }
        }

        return true;
    }

    // task3

    public static long sum(int[] arr , int n)
    {
        long cnt = 0;
        int i = 0;
        while(n-- > 0) cnt += arr[i++];
        return cnt;
    }

    public static int task3(int n, int m, int[] arr)
    {
        long sum = sum(arr, n);
        int[] prefix = new int[m+1];
        for(int i = 1; i <= n; ++i)
        {
            prefix[i] += (prefix[i-1]+arr[i-1]);
        }
        return 0;
    }

    // task4

    // task5

    // task6

    public static void main(String[] args) {
        Map<String , PriorityQueue<String>> map = new HashMap<>(Comparator.comparing());
        map.put()
    }
}