#include <iostream>
#include "model.h"

int main() {
    IloEnv env;
	IloCplex cplex(env);
	cplex.setOut(env.getNullStream());
	IloModel model(env);

	IloNumVarArray x(env);

	x = IloNumVarArray(env, 3, 0.0, 1.0, ILOFLOAT);

    populateModel(env, model,x);
    optimizeModel(cplex,env,model,x,50);

    cplex.end();
    env.end();
    return 0;
}