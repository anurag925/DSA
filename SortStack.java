import java.util.Stack;

class SortStack {
    public Stack<Integer> sort(Stack<Integer> s) {
        if(s.empty()){
            return s;
        }
        s.add(null);
        int e = s.pop();
        s = sort(s);
        s = insert(s, e);
        System.out.println(s);
        return s;
    }
    
    
    public Stack<Integer> insert(Stack<Integer> s, int e){
        System.out.println("s");
        System.out.println(s);
        System.out.println(e);
        if(s.empty()){
            return s;
        }
        if(e >= s.peek()){
            s.push(e);
            return s;
        }else{
            int temp = s.pop();
            s.push(e);
            return insert(s, temp);
        }
    }
}