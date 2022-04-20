#include <vector>
#include <ilcplex/ilocplex.h>

typedef IloArray<IloNumVarArray> NumVarMatrix;
typedef IloArray<NumVarMatrix> NumVar3Matrix;

void populateModel(IloEnv& env, IloModel& model, IloNumVarArray & x);
void optimizeModel(IloCplex & cplex, IloEnv& env, IloModel& model, IloNumVarArray & x, double total_time_limit);