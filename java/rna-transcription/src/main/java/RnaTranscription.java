class RnaTranscription {

    String transcribe(String dnaStrand) {
       char[] dnaArray = dnaStrand.toCharArray();
        for(int x =0; x < dnaArray.length; x++) {
            switch (dnaArray[x]) {
                case 'G':
                    dnaArray[x] = 'C';
                    break;
                case 'C':
                    dnaArray[x] = 'G';
                    break;
                case 'T':
                    dnaArray[x] = 'A';
                    break;
                case 'A':
                    dnaArray[x] = 'U';
                    break;
            }

       }
        return new String(dnaArray);
    }

}
