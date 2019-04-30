import java.util.HashMap;

class Scrabble {
    private static HashMap<String,Integer> scoreMap = new HashMap<String, Integer>(26) {{
        put("AEIOULNRST",1);
        put("DG",2);
        put("BCMP",3);
        put("FHVWY", 4);
        put("K",5);
        put("JX",8);
        put("QZ" ,10);
    }};
    private char[] inputWord;

    Scrabble(String word) {
        inputWord = word.toUpperCase().toCharArray();
    }

    int getScore() {
        int sum = 0;
        for(char x: inputWord) {
            for(String letters: scoreMap.keySet()){
                if(letters.contains(String.valueOf(x))) {
                    sum += scoreMap.get(letters);
                }
            }
        }
        return sum;
    }

}
