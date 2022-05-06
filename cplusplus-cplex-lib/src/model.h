#include <ilcplex/ilocplex.h>

#ifdef __cplusplus
extern "C" {
#endif
    const char * buildAndRunModel(const char * input);
#ifdef __cplusplus
}
#endif

void populateModel(IloEnv& env, IloModel& model, IloNumVarArray & x);
void optimizeModel(IloCplex & cplex, IloEnv& env, IloModel& model, IloNumVarArray & x, double total_time_limit);