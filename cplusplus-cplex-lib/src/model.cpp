#include "model.h"

void populateModel(IloEnv& env, IloModel& model, IloNumVarArray & x)
{
    model.add(x[0] + x[1] + x[2] <= 2);

    IloExpr obj(env);
    obj = operator*(1,x[0]) + operator*(1,x[1]) + operator*(1,x[2]);

    model.add(IloMaximize(env,obj));
}

void optimizeModel(IloCplex& cplex, IloEnv& env, IloModel& model, IloNumVarArray& x, double total_time_limit)
{
  cplex.extract(model);
  cplex.setParam(IloCplex::TiLim, total_time_limit);

  // Optimize the problem and obtain solution.
  if (!cplex.solve())
  {
      std::cout << "Erro na resolução do modelo" << std::endl;
      return;
  }

  IloNum curr_bound = cplex.getObjValue();
  IloNumArray optimalValues(env);
  std::cout << cplex.getCplexStatus() << ": " << curr_bound << std::endl;
  cplex.getValues(optimalValues,x);
  std::cout << "x[0] :" << optimalValues[0] << std::endl;
  std::cout << "x[1] :" << optimalValues[1] << std::endl;
  std::cout << "x[2] :" << optimalValues[2] << std::endl;
}