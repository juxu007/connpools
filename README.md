# thrift connection pools
management of multiple thrift connection pools

1. use zookeeper as registration authority; 
2. use thrift as connection protocol;
3. use weighted round robin as dynamic load balancing(DLB) strategy;
4. increase node weight when suceess, decrease node weight when failed;
5. initialization params can be setted;
