import java.io.*;
import java.util.*;

public class Main {
    public static long task1() throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        long n = Long.parseLong(br.readLine());
        return ((100+n)*(n-99)/2);
    }
    public static long task2() throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        long n = Long.parseLong(br.readLine());
        if(n == 1) return 1;
        else if(n%2==1) return (n-1)*2;
        return n*2;
    }

    public static long task3() throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());
        int[][] m = new int[n][n];
        long[][] memoization = new long[n][n];
        for(int i=0; i<n; i++){
            StringTokenizer st = new StringTokenizer(br.readLine());
            for (int j = 0; j < n; j++) {
                int num = Integer.parseInt(st.nextToken());
                m[i][j] = num;
                memoization[i][j] = -1;
            }
        }
        for(int i=0; i<n; i++){
            long sum = 0;
            for(int j=0; j<n; j++){
                sum += m[i][j];
            }
            for(int j=0; j<n; j++){
                memoization[i][j] = sum;
            }
        }
        long ans = 0;
        for(int i=0; i<n; i++){
            long sum = 0;
            for(int j=0; j<n; j++){
                sum += m[j][i];
            }
            for(int j=0; j<n; j++){
                if(Math.abs(sum-memoization[j][i]) <= m[j][i]) {
                    ++ans;
                }
            }
        }
        return ans;
    }

    public static long task4() throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(br.readLine());
        List<List<Integer>> dependencies = new ArrayList<>(n + 1);
        for (int i = 0; i <= n; i++) {
            dependencies.add(new ArrayList<>());
        }
        int[] indegree = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int a_i = Integer.parseInt(st.nextToken());
            for (int j = 0; j < a_i; j++) {
                int dep = Integer.parseInt(st.nextToken());
                dependencies.get(dep).add(i);
                indegree[i]++;
            }
        }
        return findMinCompletionTime(n, dependencies, indegree);
    }

    private static long findMinCompletionTime(int n, List<List<Integer>> dependencies, int[] indegree) {
        int[] completionTime = new int[n + 1];
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 1; i <= n; i++) {
            if (indegree[i] == 0) {
                queue.offer(i);
                completionTime[i] = 1;
            }
        }
        while (!queue.isEmpty()) {
            int process = queue.poll();
            for (int neighbor : dependencies.get(process)) {
                indegree[neighbor]--;
                completionTime[neighbor] = Math.max(completionTime[neighbor], completionTime[process] + 1);
                if (indegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        int maxCompletionTime = 0;
        for (int i = 1; i <= n; i++) {
            maxCompletionTime = Math.max(maxCompletionTime, completionTime[i]);
        }
        return maxCompletionTime;
    }

    public static void task5() throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        BufferedWriter bw = new BufferedWriter(new OutputStreamWriter(System.out));

        int n = Integer.parseInt(br.readLine());
        List<List<Integer>> dependencies = new ArrayList<>(n + 1);
        for (int i = 0; i <= n; i++) {
            dependencies.add(new ArrayList<>());
        }

        int[] indegree = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            StringTokenizer st = new StringTokenizer(br.readLine());
            int a_i = Integer.parseInt(st.nextToken());
            for (int j = 0; j < a_i; j++) {
                int dep = Integer.parseInt(st.nextToken());
                dependencies.get(dep).add(i);
                indegree[i]++;
            }
        }

        List<List<Integer>> result = findProcessLevels(n, dependencies, indegree);

        bw.write(result.size() + "\n");
        for (List<Integer> level : result) {
            Collections.sort(level);
            bw.write(level.size() + " ");
            for (int process : level) {
                bw.write(process + " ");
            }
            bw.write("\n");
        }
        bw.flush();
    }

    private static List<List<Integer>> findProcessLevels(int n, List<List<Integer>> dependencies, int[] indegree) {
        List<List<Integer>> levels = new ArrayList<>();
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 1; i <= n; i++) {
            if (indegree[i] == 0) {
                queue.offer(i);
            }
        }
        while (!queue.isEmpty()) {
            List<Integer> currentLevel = new ArrayList<>();
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                int process = queue.poll();
                currentLevel.add(process);
                for (int neighbor : dependencies.get(process)) {
                    indegree[neighbor]--;
                    if (indegree[neighbor] == 0) {
                        queue.offer(neighbor);
                    }
                }
            }
            levels.add(currentLevel);
        }
        return levels;
    }

    public static void main(String[] args) throws IOException {
        System.out.println(task2());
    }
}