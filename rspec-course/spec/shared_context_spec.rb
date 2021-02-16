RSpec.shared_context 'common' do
  before do
    @foods = []
  end
  def some_helper_method
    5
  end
  let(:some_variable) { [1, 2, 3] }
end

RSpec.describe 'first_example group' do
  include_context 'common'

  it 'can use outside instance variables' do
    expect(@foods.length).to eq(0)
    @foods << 'Sushi'
    expect(@foods.length).to eq(1)
  end

  it 'can reuse instance variables accross different examples' do
    expect(@foods.length).to eq(0)
  end
  it 'can use helper methods' do
    expect(some_helper_method).to eq(5)
  end
end

RSpec.describe 'second exmple group' do
  include_context 'common'
  it 'can use some shared let variables' do
    expect(some_variable).to eq([1, 2, 3])
  end
end

RSpec.describe Array do
  subject do
    [1, 2]
  end

  it 'is expected to equal itself' do
    expect(subject).to eq([1, 2])
  end

  it { is_expected.to eq([1, 2]) }
end
