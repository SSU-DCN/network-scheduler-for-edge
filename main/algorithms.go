package main

import (
	"fmt"
	k8sApi "k8s.io/api/core/v1"
	"log"
	"strconv"
)

// selectNode
func selectNode(nodes *k8sApi.NodeList, pod *k8sApi.Pod, scheduler Scheduler) ([]k8sApi.Node, error) {

	if len(nodes.Items) == 0 {
		return nil, fmt.Errorf("no nodes were provided")
	}

	// extract information from Pod - label values
	//appName := getDesiredFromLabels(pod, "app")
	//rtt := getDesiredFromLabels(pod, "rtt")
	bandwidth := getDesiredFromLabels(pod, "bandwidth")
	//chainPosString := getDesiredFromLabels(pod, "chainPosition")
	//nsh := getDesiredFromLabels(pod, "networkServiceHeader")
	//totalChain := getDesiredFromLabels(pod, "totalChainServ")
	//policy := getDesiredFromLabels(pod, "policy")

	//bandwidth = strings.TrimRight(minBandwidth, "Mi")
	//chainPosString = strings.TrimRight(chainPosString, "pos")
	//totalChain = strings.TrimRight(totalChain, "serv")

	podMinBandwidth := stringtoFloatBandwidth(bandwidth)
	//chainPos := stringtoInt(chainPosString)
	//totalChainServ := stringtoInt(totalChain)

	//nextApp := ""
	//prevApp := ""
	//var appList []string

	// find next and previous services in the service chain
	//if chainPos == 1 {
	//	nextApp = getDesiredFromLabels(pod, "nextService")
	//	appList = []string{nextApp}
	//} else if chainPos == totalChainServ {
	//	prevApp = getDesiredFromLabels(pod, "prevService")
	//	appList = []string{prevApp}
	//} else {
	//	prevApp = getDesiredFromLabels(pod, "prevService")
	//	nextApp = getDesiredFromLabels(pod, "nextService")
	//	appList = []string{prevApp, nextApp}
	//}

	log.Printf("Pod Name: %v \n", pod.Name)
	//log.Printf("Pod Desired location: %v \n", targetLocation)
	log.Printf("Pod Desired bandwidth: %v (Mi)\n", podMinBandwidth)
	//log.Printf("Scheduling Policy: %v \n", policy)
	//log.Printf("prevApp: %v \n", prevApp)
	//log.Printf("nextApp: %v \n", nextApp)
	//log.Printf("Service Chain: %v \n", appList)

	log.Printf("--------------------------------------------------------------\n")
	nodeList := getNodeList(nodes, podMinBandwidth)
	_, node := getMinRtt(nodeList)
	if node.GetName() == "" { // No suitable node found
		return nil, fmt.Errorf("no suitable node for target Location with enough bandwidth")
	} else {
		// update Link bandwidth
		nodeBand := getBandwidthValue(&node, "bandwidth")
		value := nodeBand - podMinBandwidth
		label := strconv.FormatFloat(value, 'f', 2, 64)
		err := updateBandwidthLabel(label, scheduler.clientset, scheduler.nodeLister, &node)
		if err != nil {
			log.Printf("encountered error when updating Bandiwdth label: %v", err)
		}
		return []k8sApi.Node{node}, nil
	}
}
