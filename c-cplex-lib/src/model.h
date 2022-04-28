/* Bring in the CPLEX function declarations and the C library
   header file stdio.h with the include of cplex.h. */

#include <ilcplex/cplex.h>

/* Bring in the declarations for the string functions */

#include <string.h>
#include <stdlib.h>

/* Include declaration for function at end of program */

static int
   setproblemdata (char **probname_p, int *numcols_p, int *numrows_p,
                   int *objsen_p, double **obj_p, double **rhs_p,
                   char **sense_p, int **matbeg_p, int **matcnt_p,
                   int **matind_p, double **matval_p,
                   double **lb_p, double **ub_p, char **ctype_p);

static void
   free_and_null (char **ptr);

   int
run_model ();