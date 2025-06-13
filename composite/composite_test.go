package composite

import (
	"testing"
)

func TestNeuron_ConnectTo(t *testing.T) {
	n1 := &Neuron{}
	n2 := &Neuron{}

	n1.ConnectTo(n2)

	if len(n1.Out) != 1 || n1.Out[0] != n2 {
		t.Errorf("n1.Out should contain n2")
	}
	if len(n2.In) != 1 || n2.In[0] != n1 {
		t.Errorf("n2.In should contain n1")
	}
}

func TestNeuronLayer_Iter(t *testing.T) {
	layer := NewNeuronLayer(3)
	neurons := layer.Iter()
	if len(neurons) != 3 {
		t.Errorf("Expected 3 neurons, got %d", len(neurons))
	}
}

func TestConnect_NeuronAndLayer(t *testing.T) {
	n := &Neuron{}
	layer := NewNeuronLayer(2)
	Connect(n, layer)
	if len(n.Out) != 2 {
		t.Errorf("Expected neuron to connect to 2 neurons, got %d", len(n.Out))
	}
	for i, out := range n.Out {
		if out != &layer.Neurons[i] {
			t.Errorf("Expected out neuron %d to be layer neuron", i)
		}
		if len(out.In) != 1 || out.In[0] != n {
			t.Errorf("Expected layer neuron %d to be connected from n", i)
		}
	}
}

func TestConnect_LayerToLayer(t *testing.T) {
	layer1 := NewNeuronLayer(2)
	layer2 := NewNeuronLayer(3)
	Connect(layer1, layer2)
	for i, n := range layer1.Neurons {
		if len(n.Out) != 3 {
			t.Errorf("Neuron %d in layer1 should have 3 outputs, got %d", i, len(n.Out))
		}
		for j, out := range n.Out {
			if out != &layer2.Neurons[j] {
				t.Errorf("Expected neuron %d in layer1 to connect to neuron %d in layer2", i, j)
			}
		}
	}
	for j, n := range layer2.Neurons {
		if len(n.In) != 2 {
			t.Errorf("Neuron %d in layer2 should have 2 inputs, got %d", j, len(n.In))
		}
	}
}
