#include<stdio.h>

int anzahl;

void fill (int *array){
    printf("\ninsert the numbers for sorting and seperate them with backspackes ...\n");
    for (int i=0;i<anzahl;i++){
	printf("Zahl %2i: ",i+1);
	scanf("%d",&array[i]);
    }
    printf("\n");
}

void sort(int *array){
	/*gehe das array durch und tausche jede zahl solange wie die anzahl der kleineren zahlen im array nicht dem index der zahl im array+1 entspricht*/
	int i=0;
	while (i<anzahl){
		int lower=0;
		for (int j=0;j<anzahl;j++){
			if (array[j]<array[i]){
				lower++;
			}
		}
		
		/*swap*/
		int tmp=array[i];
		array[i]=array[lower+1];
		array[lower+1]=tmp;
		
		if (i==lower+1){
			i++;
		}
	}
	
}

void print(int *array){
	printf("arraystart: ");
	for (int i=0;i<anzahl;i++){
		printf("%8d",array[i]);
	}
	printf(" :arrayend\n");
}

int main(){
	/*start and user input of numbers for sorting*/
	printf("***************************** Swap Sort *****************************");
	
	printf("\n\nhow many numbers do you want to sort? ");
	scanf("%d",&anzahl);
	
	int array[anzahl];
	fill(array);
	
	
	print(array);
	
	/*START of sorting array*/
	int i=0;
	
	while (i<anzahl){
		/*test whether is already sorted*/
		int unsorted=0;
		for (int j=1;j<anzahl;j++){
			if (array[j]<array[j-1]){
				unsorted++;
			}
		}
		if (unsorted==0){
			break;
		}
		
		/*counting of the lower than insspected numbers*/
		int lower=0;
		for (int j=0;j<anzahl;j++){
			if (array[j]<array[i]){
				lower++;
			}
		}

		/*swap*/
		int tmp=array[i];
		array[i]=array[lower];
		array[lower]=tmp;

		
		if (i==lower){
			i++;
			printf("\ni++\t");
		} else {
			printf("\t");
		}
		
		print(array);
		printf("\n");
		
	}
	/*END of sorting array*/
		
	return 0;
}
