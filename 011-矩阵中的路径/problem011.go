package main

// class Solution {
//     public boolean exist(char[][] board, String word) {
//         char[] words = word.toCharArray();
//         for(int i = 0; i < board.length; i++) {
//             for(int j = 0; j < board[0].length; j++) {
//                 if(dfs(board, words, i, j, 0)) return true;
//             }
//         }
//         return false;
//     }
//     boolean dfs(char[][] board, char[] word, int i, int j, int k) {
//         if(i >= board.length || i < 0 || j >= board[0].length || j < 0 || board[i][j] != word[k]) return false;
//         if(k == word.length - 1) return true;
//         board[i][j] = '\0';
//         boolean res = dfs(board, word, i + 1, j, k + 1) || dfs(board, word, i - 1, j, k + 1) ||
//                       dfs(board, word, i, j + 1, k + 1) || dfs(board, word, i , j - 1, k + 1);
//         board[i][j] = word[k];
//         return res;
//     }
// }

func exist(board [][]byte, word string) bool {
	for i := 0;i < len(board);i++{
		for j := 0;j < len(board[0]);j++{
			if dfs(board,word,i,j,0) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, word string, i,j,k int) bool {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
		return false
	}
	if k == len(word) - 1 {
		return true
	}
	board[i][j] = ' '
	res := dfs(board,word,i+1,j,k+1) || dfs(board,word,i-1,j,k+1) ||
		dfs(board,word,i,j+1,k+1) || dfs(board,word,i,j-1,k+1)
	board[i][j] = word[k]
	return res
}
