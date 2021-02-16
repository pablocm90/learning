RSpec.describe 'Allow method review' do
  it 'can return value for methods on doubles' do
    calculator = double
    allow(calculator).to receive(:add).and_return(15)
    expect(calculator.add).to eq(15)
  end
  it 'can stub one or more methods on a real object' do
    arr = [1, 2, 3]
    expect(arr.sum).to eq(6)
    allow(arr).to receive(:sum).and_return(10)
    expect(arr.sum).to eq(10)
  end
  it 'can return multiple return values in sequence' do
    mock_array = %i[b c]
    allow(mock_array).to receive(:pop).and_return(:c, :b, nil)
    expect(mock_array.pop).to eq(:c)
    expect(mock_array.pop).to eq(:b)
    expect(mock_array.pop).to be_nil
  end
end
