RSpec.describe 'All Matcher' do 
  it 'allows for aggregate checks' do 
    expect([5, 7, 9, 13]).to all(be_odd)
    expect([5, 7, 9]).to all(be < 10)

  end

  describe [5, 7, 9] do
    it { is_expected.to all(be_odd) }
  end
end
